package main

import (
	"main/internal/users"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	users.RegisterRoutes(router, users.New())

	router.Run(":3000")
}
