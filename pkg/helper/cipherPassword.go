package helper

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"key-haven-back/pkg/secret"
	"strings"
)

func EncryptPassword(plaintextPassword string, masterPassword string) (string, error) {
	if plaintextPassword == "" {
		return "", fmt.Errorf("password to be encrypted cannot be empty")
	}
	if masterPassword == "" {
		return "", fmt.Errorf("master password cannot be empty")
	}

	masterPasswordMd5 := md5.Sum([]byte(masterPassword))
	key := masterPasswordMd5[:]

	ciphertext, nonce, err := secret.EncryptAES(plaintextPassword, key)
	if err != nil {
		return "", fmt.Errorf("encryption error: %v", err)
	}

	encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)
	encodedNonce := base64.StdEncoding.EncodeToString(nonce)

	return fmt.Sprintf("%s:%s", encodedCiphertext, encodedNonce), nil
}

func DecryptPassword(ciphertextPasswordNonce string, masterPassword string) (string, error) {
	if ciphertextPasswordNonce == "" {
		return "", fmt.Errorf("encrypted text cannot be empty")
	}
	if masterPassword == "" {
		return "", fmt.Errorf("master password cannot be empty")
	}

	parts := strings.Split(ciphertextPasswordNonce, ":")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid format: expected 'ciphertext:nonce'")
	}

	if parts[0] == "" {
		return "", fmt.Errorf("ciphertext cannot be empty")
	}
	if parts[1] == "" {
		return "", fmt.Errorf("nonce cannot be empty")
	}

	ciphertext, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return "", fmt.Errorf("failed to decode ciphertext: %v", err)
	}

	nonce, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", fmt.Errorf("failed to decode nonce: %v", err)
	}

	masterPasswordMd5 := md5.Sum([]byte(masterPassword))
	key := masterPasswordMd5[:]

	decrypted, err := secret.DecryptAES(ciphertext, key, nonce)
	if err != nil {
		return "", fmt.Errorf("decryption error: %v", err)
	}

	return decrypted, nil
}
