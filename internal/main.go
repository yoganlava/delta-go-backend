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
	stripeService := stripes.New()
	authService := auth.New()
	creatorsService := creators.New()
	postsService := posts.New()
	projectsService := projects.New()
	usersService := users.New()
	stripes.RegisterRoutes(router, &stripeService, &creatorsService, &postsService)
	return
	auth.RegisterRoutes(router, authService)
	creators.RegisterRoutes(router, &creatorsService)
	posts.RegisterRoutes(router, &postsService)
	projects.RegisterRoutes(router, &projectsService)
	users.RegisterRoutes(router, &usersService)
	router.Run(":5000")
}
