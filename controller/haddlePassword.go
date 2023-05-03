package controller

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) // Salt rounds = 14 use 1.5 sec to hash password
	return string(bytes), err
}

func CheckPassword(password string, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err == nil {
		return true
	} else {
		return false
	}
}
