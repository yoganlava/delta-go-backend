package middleware

import (
	"fmt"
	"main/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

//JwtMiddleware used in places where non-users cannot go
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerToken = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < 1 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		tokenString := authHeader[len(bearerToken):]
		// Might want to remove New() everytime we want to verify token as there is no need
		token, err := auth.New().VerifyToken(tokenString)

		if token > 0 {
			c.Set("user_id", token)
		} else {
			fmt.Print(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

//OptionalMiddleware used in places where user id is not needed but used
func OptionalMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) > 0 {
			tokenString := authHeader[len("Bearer"):]
			// Might want to remove New() everytime we want to verify token as there is no need
			token, err := auth.New().VerifyToken(tokenString)
			if err != nil {
				fmt.Print(err.Error())
			} else {
				c.Set("user_id", token)
			}
		}
	}
}
