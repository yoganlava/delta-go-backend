package messaging

import (
	"main/internal/dto"
	"main/internal/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service MessagingService
}

func RegisterRoutes(router *gin.Engine, service MessagingService) {
	// c := controller{service}
	// r := router.Group("/messages")
	// r.Use(middleware.JwtMiddleware()) {

	// }
}

func (con controller) SendMessage(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "Not logged in",
		})
		return
	}
	var err error
	defer utility.ErrorHandleHTTP(c, err)
	var m dto.CreateMessageDTO
	c.BindQuery(&m)
	m.SenderID = userID.(int)
	err = con.service.SendMessage(m)
	if err != nil {
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Success",
	})
}

func (con controller) RetrieveSentMessages(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "Not logged in",
		})
		return
	}
	var err error
	defer utility.ErrorHandleHTTP(c, err)
	messages, err := con.service.RetrieveUserSentMessages(userID.(int))
	if err != nil {
		return
	}
	c.JSON(http.StatusAccepted, messages)
}

func (con controller) RetrieveUserMessages(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "Not logged in",
		})
		return
	}
	var err error
	defer utility.ErrorHandleHTTP(c, err)
	messages, err := con.service.RetrieveUserMessages(userID.(int))
	if err != nil {
		return
	}
	c.JSON(http.StatusAccepted, messages)
}
