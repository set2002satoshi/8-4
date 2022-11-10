package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparisonPassAndHash(a, b string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(a), []byte(b)); err != nil {
		return false
	}
	return true
}
