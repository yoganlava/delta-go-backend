package main

import (
	"main/internal/auth"
	"main/internal/creators"
	"main/internal/posts"
	"main/internal/users"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := gin.Default()
	auth.RegisterRoutes(router, auth.New())
	creators.RegisterRoutes(router, creators.New())
	posts.RegisterRoutes(router, posts.New())
	users.RegisterRoutes(router, users.New())
	router.Run(":5000")
}
