package users

import (
	"main/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service UserService
}

// RegisterRoutes register routes
func RegisterRoutes(router *gin.Engine, service UserService) {
	c := controller{service}
	u := router.Group("/users")
	u.Use(middleware.JwtMiddleware())
	{
		u.GET("/me", c.fetchSelf)
	}

	// u.POST("/login", c.login)
}

// func (con controller) register(c *gin.Context) {
// 	var request RegisterRequest
// 	if err := c.BindJSON(&request); err != nil {
// 		c.AbortWithStatus(400)
// 		return
// 	}

// 	c.JSON(200, con.service.Register(request))
// }

func (con controller) fetchSelf(c *gin.Context) {
	id, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// var user entity.SelfUser
	user, err := con.service.FetchSelf(id.(int))
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"user": user,
	})
	return
}
