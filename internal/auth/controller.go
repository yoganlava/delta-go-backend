package auth

import (
	"github.com/gin-gonic/gin"
)

type controller struct {
	service AuthService
}

func RegisterRoutes(router *gin.Engine, service AuthService) {
	c := controller{service}
	u := router.Group("/auth")
	u.POST("/register", c.register)
	u.POST("/login", c.login)
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
