package admin

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(raw string) (string, error) {
	password := strings.TrimSpace(raw)
	if password == "" {
		return "", fmt.Errorf("password is required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func verifyPassword(hash string, password string) bool {
	if strings.TrimSpace(hash) == "" || strings.TrimSpace(password) == "" {
		return false
	}
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
