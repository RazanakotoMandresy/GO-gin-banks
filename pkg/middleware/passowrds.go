package middleware

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// handle two func of bcrypt
func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Printf("Erreur lors du cryptage du mots de passe : %v \n", err)
		return ""
	}
	return string(bytes)
}
func IsTruePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
