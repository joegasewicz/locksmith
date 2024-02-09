package utilities

import "golang.org/x/crypto/bcrypt"

// Hash password
func Hash(password string) (string, error) {
	var err error
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

// Compare hash with password
func Compare(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
