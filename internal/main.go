package main

import (
	"main/internal/auth"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := gin.Default()
	auth.RegisterRoutes(router, auth.New())

	router.Run(":5000")
}
