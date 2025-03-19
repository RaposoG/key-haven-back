package helper

import (
	"crypto/md5"
	"fmt"
	"key-haven-back/pkg/secret"
	"strings"
)

func EncryptPassword(plaintextPassword string, masterPassword string) (string, error) {
	masterPasswordMd5 := md5.Sum([]byte(masterPassword))
	key := masterPasswordMd5[:]
	ciphertext, nonce, err := secret.EncryptAES(plaintextPassword, key)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%s", ciphertext, nonce), nil
}

func DecryptPassword(ciphertextPasswordNounce string, masterPassword string) (string, error) {
	masterPasswordMd5 := md5.Sum([]byte(masterPassword))
	key := masterPasswordMd5[:]
	splited := strings.Split(ciphertextPasswordNounce, ":")
	decrypted, err := secret.DecryptAES([]byte(splited[0]), key, []byte(splited[1]))
	if err != nil {
		return "", err
	}
	return decrypted, nil
}
