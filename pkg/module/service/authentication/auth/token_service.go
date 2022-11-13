package auth

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// issue login token
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

func CheckLoggedIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := godotenv.Load(".env")
		if err != nil {
			ctx.JSON(401, "Unauthorized")
			ctx.Abort()
		}
		
		clientKey := ctx.Request.Header.Get("clientKey")

		token, err := jwt.ParseWithClaims(clientKey, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil {
			ctx.JSON(401, "Unauthorized")
			ctx.Abort()
		}

		claims := token.Claims.(*jwt.StandardClaims)
		ctx.Set("userID", claims.Issuer)

		ctx.Next()
	}
}