package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CheckHashPassword(hashpassword string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password)); err != nil {
		return fmt.Errorf("invalid password")
	}
	return nil
}
