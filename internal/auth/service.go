package auth

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserClaim struct {
	jwt.MapClaims
	id  int
	exp int64
}

// CreateToken for user
func CreateToken(id int) string {
	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"id":  id,
	// 	"exp": time.Now().Add(time.Minute * 30).Unix(),
	// })
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		id:  id,
		exp: time.Now().Add(time.Minute * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		panic(err)
	}
	return tokenString
}

// VerifyToken and return user id or error
func VerifyToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Something went wrong when parsing")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return -1, err
	}
	if !token.Valid {
		return -1, errors.New("Token expired")
	}
	return token.Claims.(*UserClaim).id, nil
}

// Middleware is run ever time its group has been requested
func Middleware(c *gin.Context) {
	_, err := VerifyToken(strings.Split(c.GetHeader("Authorization"), " ")[1])
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	return
}
