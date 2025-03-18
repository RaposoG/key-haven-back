package cryptoEAS

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func encryptAES(plaintext string, key []byte) ([]byte, []byte) {
	block, _ := aes.NewCipher(key)
	nonce := make([]byte, 12)
	io.ReadFull(rand.Reader, nonce)

	aesgcm, _ := cipher.NewGCM(block)
	ciphertext := aesgcm.Seal(nil, nonce, []byte(plaintext), nil)
	return ciphertext, nonce
}

func decryptAES(ciphertext, key, nonce []byte) string {
	block, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(block)
	plaintext, _ := aesgcm.Open(nil, nonce, ciphertext, nil)
	return string(plaintext)
}

func main() {
	key := []byte("12345678901234567890123456789012") // 32 bytes para AES-256
	plaintext := "Mensagem secreta"

	ciphertext, nonce := encryptAES(plaintext, key)
	fmt.Printf("Criptografado: %x\n", ciphertext)

	decrypted := decryptAES(ciphertext, key, nonce)
	fmt.Println("Descriptografado:", decrypted)
}
