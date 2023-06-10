package security

import "golang.org/x/crypto/bcrypt"

// Hash receive a string and add a has in it
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// PasswordValidade compare a string password and a hash and returns if they are equal
func PasswordValidate(hashPassword, stringPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(stringPassword))
}