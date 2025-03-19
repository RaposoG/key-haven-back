package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
)

func EncryptAES(plaintext string, key []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

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

func main() {
	senha := []byte("lindaosafadao")
	senhaMd5 := md5.Sum(senha)
	key := senhaMd5[:]
	plaintext := "Mensagem secreta"

	ciphertext, nonce, _ := EncryptAES(plaintext, key)
	fmt.Printf("Criptografado: %x\n", ciphertext)

	decrypted, _ := DecryptAES(ciphertext, key, nonce)
	fmt.Println("Descriptografado:", decrypted)
}
