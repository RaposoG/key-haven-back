package unittests

import (
	"key-haven-back/pkg/helper"
	"strings"
	"testing"
)

func TestEncryptAndDecryptSimplePassword(t *testing.T) {
	password := "mySimplePassword123"
	masterPassword := "MasterKey123!"

	encrypted, err := helper.EncryptPassword(password, masterPassword)
	if err != nil {
		t.Fatalf("Failed to encrypt password: %v", err)
	}

	if encrypted == password {
		t.Errorf("Encrypted password is identical to original password")
	}

	if !strings.Contains(encrypted, ":") {
		t.Errorf("Encrypted string does not contain separator")
	}

	parts := strings.Split(encrypted, ":")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		t.Errorf("Invalid encrypted format: %s", encrypted)
	}

	decrypted, err := helper.DecryptPassword(encrypted, masterPassword)
	if err != nil {
		t.Fatalf("Failed to decrypt password: %v", err)
	}

	if decrypted != password {
		t.Errorf("Decrypted password does not match original password. Got %s, expected %s", decrypted, password)
	}
}

func TestEncryptWithEmptyPassword(t *testing.T) {
	password := ""
	masterPassword := "MasterKey123!"

	_, err := helper.EncryptPassword(password, masterPassword)
	if err == nil {
		t.Errorf("Expected an error when encrypting with empty password, but no error was returned")
	}
}

func TestEncryptWithEmptyMasterPassword(t *testing.T) {
	password := "securePassword123"
	masterPassword := ""

	_, err := helper.EncryptPassword(password, masterPassword)
	if err == nil {
		t.Errorf("Expected an error when encrypting with empty master password, but no error was returned")
	}
}

func TestDecryptWithEmptyCiphertext(t *testing.T) {
	ciphertext := ""
	masterPassword := "MasterKey123!"

	_, err := helper.DecryptPassword(ciphertext, masterPassword)
	if err == nil {
		t.Errorf("Expected an error when decrypting with empty ciphertext, but no error was returned")
	}
}

func TestDecryptWithEmptyMasterPassword(t *testing.T) {

	originalPassword := "securePassword123"
	originalMasterPassword := "MasterKey123!"

	encrypted, err := helper.EncryptPassword(originalPassword, originalMasterPassword)
	if err != nil {
		t.Fatalf("Failed to create encrypted text for test: %v", err)
	}

	emptyMasterPassword := ""

	_, err = helper.DecryptPassword(encrypted, emptyMasterPassword)
	if err == nil {
		t.Errorf("Expected an error when decrypting with empty master password, but no error was returned")
	}
}

func TestEncryptAndDecryptSpecialChars(t *testing.T) {
	password := "P@ssw0rd!#$%^&*()"
	masterPassword := "M@st3r!#$%^&*()"

	encrypted, err := helper.EncryptPassword(password, masterPassword)
	if err != nil {
		t.Fatalf("Failed to encrypt password with special characters: %v", err)
	}

	if encrypted == password {
		t.Errorf("Encrypted password is identical to original password with special characters")
	}

	decrypted, err := helper.DecryptPassword(encrypted, masterPassword)
	if err != nil {
		t.Fatalf("Failed to decrypt password with special characters: %v", err)
	}

	if decrypted != password {
		t.Errorf("Decrypted password does not match original password with special characters. Got %s, expected %s", decrypted, password)
	}
}

func TestEncryptAndDecryptLongPassword(t *testing.T) {
	password := "ThisIsAVeryLongPasswordThatExceedsFiftyCharactersInLengthToTestBoundaries"
	masterPassword := "MasterKey123!"

	encrypted, err := helper.EncryptPassword(password, masterPassword)
	if err != nil {
		t.Fatalf("Failed to encrypt long password: %v", err)
	}

	if encrypted == password {
		t.Errorf("Encrypted password is identical to original long password")
	}

	decrypted, err := helper.DecryptPassword(encrypted, masterPassword)
	if err != nil {
		t.Fatalf("Failed to decrypt long password: %v", err)
	}

	if decrypted != password {
		t.Errorf("Decrypted password does not match original long password. Got %s, expected %s", decrypted, password)
	}
}

func TestDecryptWithWrongMasterPassword(t *testing.T) {
	originalPassword := "mySecurePassword123"
	correctMasterPassword := "CorrectMasterKey!"
	wrongMasterPassword := "WrongMasterKey!"

	encrypted, err := helper.EncryptPassword(originalPassword, correctMasterPassword)
	if err != nil {
		t.Fatalf("Failed to encrypt password: %v", err)
	}

	_, err = helper.DecryptPassword(encrypted, wrongMasterPassword)

	if err == nil {
		t.Error("Expected an error when decrypting with wrong master password, but no error was returned")
	}
}

func TestDecryptMissingSeparator(t *testing.T) {
	invalidCiphertext := "invalidseparator"
	masterPassword := "MasterKey123!"

	_, err := helper.DecryptPassword(invalidCiphertext, masterPassword)
	if err == nil {
		t.Errorf("Expected an error with invalid format (missing separator), but no error was returned")
	}
}

func TestDecryptEmptyCiphertext(t *testing.T) {
	invalidCiphertext := ":nonce"
	masterPassword := "MasterKey123!"

	_, err := helper.DecryptPassword(invalidCiphertext, masterPassword)
	if err == nil {
		t.Errorf("Expected an error with empty ciphertext, but no error was returned")
	}
}

func TestDecryptEmptyNonce(t *testing.T) {
	invalidCiphertext := "ciphertext:"
	masterPassword := "MasterKey123!"

	_, err := helper.DecryptPassword(invalidCiphertext, masterPassword)
	if err == nil {
		t.Errorf("Expected an error with empty nonce, but no error was returned")
	}
}

func TestDecryptInvalidBase64Ciphertext(t *testing.T) {
	invalidCiphertext := "not-valid-base64:validbase64=="
	masterPassword := "MasterKey123!"

	_, err := helper.DecryptPassword(invalidCiphertext, masterPassword)
	if err == nil {
		t.Errorf("Expected an error with invalid base64 in ciphertext, but no error was returned")
	}
}

func TestDecryptInvalidBase64Nonce(t *testing.T) {
	invalidCiphertext := "validbase64==:not-valid-base64"
	masterPassword := "MasterKey123!"

	_, err := helper.DecryptPassword(invalidCiphertext, masterPassword)
	if err == nil {
		t.Errorf("Expected an error with invalid base64 in nonce, but no error was returned")
	}
}
