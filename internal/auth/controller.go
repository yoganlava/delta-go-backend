package auth

import (
	"main/internal/dto"
	"net/http"

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
	var request dto.AuthRegister
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatus(500)
		return
	}

	err := con.service.Register(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}

func (con controller) login(c *gin.Context) {
	var request dto.AuthLogin
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatus(500)
		return
	}
	user, err := con.service.Login(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else if user.ID != 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "ユーザー名・Eメールとパスワードの組み合わせのアカウントが見つかりませんでした",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"user": user,
	})
}
