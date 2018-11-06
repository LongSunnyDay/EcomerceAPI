package user

import (
	"golang.org/x/crypto/bcrypt"
)

// Password encryption and authorization done by example:
// https://gowebexamples.com/password-hashing/

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}



