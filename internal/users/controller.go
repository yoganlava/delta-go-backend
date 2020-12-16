package users

import (
	"github.com/gin-gonic/gin"
)

type controller struct {
	service UserService
}

// RegisterRoutes register routes
func RegisterRoutes(router *gin.Engine, service UserService) {
	c := controller{service}
	router.POST("/register", c.register)
	router.POST("/login", c.login)
}

func (con controller) register(c *gin.Context) {
	var request RegisterRequest
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatus(400)
		return
	}

	c.JSON(200, con.service.Register(request))
}

func (con controller) login(c *gin.Context) {

}
