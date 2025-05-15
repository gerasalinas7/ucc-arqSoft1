package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashea la contraseña
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
