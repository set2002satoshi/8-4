package auth

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func IssueToken(userID int) (string, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return "", errors.New("couldn't load Secret")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(userID),
		ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
	})

	token, err := claims.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		// return "", errors.New("Incorrect secret key")
		return "", jwt.ErrInvalidKey
	}
	return token, nil
}
