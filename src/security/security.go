package security

import "golang.org/x/crypto/bcrypt"

// HASH generates a hash from a password
func HASH(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VERIFY verifies if a password and a hash are the same
func VERIFY(password string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
