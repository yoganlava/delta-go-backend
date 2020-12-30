package messaging

import (
	"main/internal/dto"
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
	var m dto.CreateMessageDTO
	c.BindQuery(&m)
	m.SenderID = userID.(int)
	err := con.service.SendMessage(m)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
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
	messages, err := con.service.RetrieveUserSentMessages(userID.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusAccepted, gin.H{
			"error": err.Error(),
		})
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
	messages, err := con.service.RetrieveUserMessages(userID.(int))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, messages)
}
