package auth

import (
	"fmt"
	"main/internal/dto"
	"main/internal/utility"
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
	var err error
	defer utility.ErrorHandleHTTP(c, err)
	if err := c.BindJSON(&request); err != nil {
		return
	}

	user, err := con.service.Register(&request)
	if err != nil {
		return
	}
	jwt := CreateToken(user.ID)
	c.JSON(http.StatusCreated, gin.H{
		"user": user,
		"jwt":  jwt,
	})
}

func (con controller) login(c *gin.Context) {
	var request dto.AuthLogin
	var err error
	defer utility.ErrorHandleHTTP(c, err)
	if err := c.BindJSON(&request); err != nil {
		return
	}
	user, err := con.service.Login(request)
	if err != nil {
		return
	} else if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "ユーザー名・Eメールとパスワードの組み合わせのアカウントが見つかりませんでした",
		})
		return
	}
	fmt.Print(user)
	jwt := CreateToken(user.ID)
	c.JSON(http.StatusAccepted, gin.H{
		"user": user,
		"jwt":  jwt,
	})
}
