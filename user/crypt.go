package user

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func getPwd() []byte {
	fmt.Println("Enter a password")
	var pwd string
	_, err := fmt.Scan(&pwd)
	if err != nil {
		log.Println(err)
	}
	return []byte(pwd)
}



