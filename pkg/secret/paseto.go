package secret

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"key-haven-back/config"
	"log"
	"strings"
	"time"

	"aidanwoods.dev/go-paseto"
)

var (
	ErrTokenExpired = errors.New("token has expired")
	ErrInvalidToken = errors.New("invalid token")
	ErrEmptyToken   = errors.New("empty token provided")
	ErrInvalidInput = errors.New("invalid input parameters")
)

type TokenClaims struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	ExpiresAt time.Time `json:"expires_at"`
}

// Generate a secure key or load from environment
func getOrCreateSymmetricKey() paseto.V4SymmetricKey {
	keyHex := config.GetEnvOrDefault("PASETO_KEY", "")
	if keyHex == "" {
		// Generate a new symmetric key if not found
		key := paseto.NewV4SymmetricKey()
		_ = hex.EncodeToString(key.ExportBytes())
		log.Println("Warning: Generated new PASETO key. In production, store this securely.")
		return key
	}

	// Decode the hex key
	keyBytes, err := hex.DecodeString(keyHex)
	if err != nil {
		log.Printf("Error decoding PASETO key: %v. Generating new key.", err)
		// Fall back to generating a new symmetric key
		key := paseto.NewV4SymmetricKey()
		return key
	}

	// Use the correct function to load a symmetric key
	key, err := paseto.V4SymmetricKeyFromBytes(keyBytes)
	if err != nil {
		log.Printf("Error creating symmetric key from bytes: %v. Generating new key.", err)
		return paseto.NewV4SymmetricKey()
	}

	return key
}

// GenerateToken creates a secure PASETO token with user information
func GenerateToken(userID, email string, expiresIn time.Duration) (string, error) {
	// Validate input
	if strings.TrimSpace(userID) == "" || strings.TrimSpace(email) == "" {
		return "", fmt.Errorf("%w: userID and email cannot be empty", ErrInvalidInput)
	}

	if expiresIn <= 0 {
		return "", fmt.Errorf("%w: expiration time must be positive", ErrInvalidInput)
	}

	key := getOrCreateSymmetricKey()
	token := paseto.NewToken()

	// Add standard claims
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(expiresIn))

	// Add custom claims
	token.SetString("user_id", userID)
	token.SetString("email", email)

	// Generate a random nonce for added security
	nonce := make([]byte, 32)
	if _, err := rand.Read(nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", err)
	}
	token.SetString("nonce", hex.EncodeToString(nonce))

	// Encrypt and sign the token
	encrypted := token.V4Encrypt(key, nil)
	return encrypted, nil
}

// ValidateToken verifies and extracts claims from a PASETO token
func ValidateToken(tokenString string) (*TokenClaims, error) {
	if tokenString == "" {
		return nil, ErrEmptyToken
	}

	key := getOrCreateSymmetricKey()
	parser := paseto.NewParser()

	// Set validation rules
	parser.AddRule(paseto.NotExpired())
	parser.AddRule(paseto.ValidAt(time.Now()))

	// Parse and validate the token
	token, err := parser.ParseV4Local(key, tokenString, nil)
	if err != nil {
		if err.Error() == "token has expired" {
			return nil, ErrTokenExpired
		}
		return nil, fmt.Errorf("%w: %v", ErrInvalidToken, err)
	}

	// Extract claims
	userID, err := token.GetString("user_id")
	if err != nil {
		return nil, fmt.Errorf("%w: missing user_id claim", ErrInvalidToken)
	}

	email, err := token.GetString("email")
	if err != nil {
		return nil, fmt.Errorf("%w: missing email claim", ErrInvalidToken)
	}

	// Extract expiration time
	expClaim, err := token.GetTime("exp")
	if err != nil {
		return nil, fmt.Errorf("%w: missing expiration claim", ErrInvalidToken)
	}

	return &TokenClaims{
		UserID:    userID,
		Email:     email,
		ExpiresAt: expClaim,
	}, nil
}

// RefreshToken creates a new token with fresh expiration time
func RefreshToken(tokenString string, newExpiresIn time.Duration) (string, error) {
	claims, err := ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	// Generate a new token with the same user data
	return GenerateToken(claims.UserID, claims.Email, newExpiresIn)
}
