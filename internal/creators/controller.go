package creators

import (
	"fmt"
	"main/internal/dto"
	"main/internal/middleware"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service CreatorService
}

// RegisterRoutes register routes
func RegisterRoutes(router *gin.Engine, service CreatorService) {
	c := controller{service}
	r := router.Group("/creators")
	r.GET("/:id", c.FetchCreator)
	r.Use(middleware.JwtMiddleware())
	{
		r.POST("/", c.CreateCreator)
	}
}

func (con controller) FetchCreator(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	creator, err := con.service.FetchCreator(id)
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{
			"error": "作成者が見つかりません",
		})
		return
	}
	c.JSON(http.StatusAccepted, creator)
}

func (con controller) CreateCreator(c *gin.Context) {
	var creator dto.CreateCreatorDTO
	if err := c.BindJSON(&creator); err != nil {
		c.AbortWithStatus(500)
		return
	}

	id, exists := c.Get("user_id")
	fmt.Print(id)
	creator.UserID = id.(int)

	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	newCreator, err := con.service.CreateCreator(creator)
	fmt.Print(err)
	if err != nil || newCreator.ID == 0 {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"creator": newCreator,
	})
}
