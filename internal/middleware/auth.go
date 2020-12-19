package middleware

import (
	"fmt"
	"main/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerToken = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(bearerToken):]
		token, err := auth.New().VerifyToken(tokenString)
		if token != -1 {
			fmt.Print(token)
			c.Set("user_id", token)
		} else {
			fmt.Print(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func OptionalMiddelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) > 0 {
			tokenString := authHeader[len("Bearer"):]
			token, err := auth.New().VerifyToken(tokenString)
			if err != nil {
				fmt.Print(err.Error())
			} else {
				c.Set("user_id", token)
			}

		}
	}
}
