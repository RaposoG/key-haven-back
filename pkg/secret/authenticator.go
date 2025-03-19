package secret

import (
	"key-haven-back/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func makeClaims(userID string, roles []string) jwt.MapClaims {
	return jwt.MapClaims{
		"sub":   userID,
		"roles": roles,
		"exp":   time.Now().Add(time.Minute * 15).Unix(),
	}
}

func CreateJwtToken(userID string, roles []string) (string, error) {
	secret := config.GetEnvOrDefault("JWT_SECRET", "secret")
	claims := makeClaims(userID, roles)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseJwtToken(tokenString string, claims jwt.Claims) error {
	secret := config.GetEnvOrDefault("JWT_SECRET", "secret")
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return jwt.ErrSignatureInvalid
	}

	return nil
}
