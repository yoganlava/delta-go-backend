package main

import (
	"main/internal/auth"
	"main/internal/creators"
	"main/internal/posts"
	"main/internal/projects"
	"main/internal/stripes"
	"main/internal/users"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v72"
)

func main() {
	godotenv.Load()
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	router := gin.Default()
	stripes.RegisterRoutes(router, stripes.New())
	auth.RegisterRoutes(router, auth.New())
	creators.RegisterRoutes(router, creators.New())
	posts.RegisterRoutes(router, posts.New())
	projects.RegisterRoutes(router, projects.New())
	users.RegisterRoutes(router, users.New())
	router.Run(":5000")
}
