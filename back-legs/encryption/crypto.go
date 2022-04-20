package encryption

import (
	bcrypt "golang.org/x/crypto/bcrypt"
)

func EncryptText(text string) (string, error) {
	hashedText, err := bcrypt.GenerateFromPassword([]byte(text), 8)
	if err != nil {
		return "err", err
	}
	return string(hashedText), nil
}

func Compare(text string, encryptedText string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(encryptedText), []byte(text)); err != nil {
		return false
	}
	return true
}

// Compare the stored hashed password, with the hashed version of the password that was received
