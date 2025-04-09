package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(hash, password string) error {
	fmt.Println(hash, password)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
