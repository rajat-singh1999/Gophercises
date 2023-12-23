package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func CreatePasswordHash(password string) (string, error) {
	pass := []byte(password)
	hashedPass, err := bcrypt.GenerateFromPassword(pass, 15)
	if err != nil {
		return "", err
	}
	return string(hashedPass), nil
}

func CheckPasswordHash(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
