package main

import (
	"main/internal/users"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := gin.Default()

	users.RegisterRoutes(router, users.New())

	router.Run(":3000")
}
