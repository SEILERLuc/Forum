package server

import (
	"golang.org/x/crypto/bcrypt"
)

// File for the users's passwords :
// - hash a user password
// - checking between a hashed password and a password

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
