package main

import (
	"main/internal/auth"
	"main/internal/users"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := gin.Default()
	auth.RegisterRoutes(router, auth.New())

	users.RegisterRoutes(router, users.New())
	router.Run(":5000")
}
