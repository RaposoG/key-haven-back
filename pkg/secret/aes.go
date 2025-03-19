package secret

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func EncryptAES(plaintext string, key []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	nonce := make([]byte, 12)
	io.ReadFull(rand.Reader, nonce)

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}
	ciphertext := aesgcm.Seal(nil, nonce, []byte(plaintext), nil)
	return ciphertext, nonce, nil
}

func DecryptAES(ciphertext, key, nonce []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
