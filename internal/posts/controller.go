package posts

import (
	"fmt"
	"main/internal/dto"
	"main/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service PostService
}

// RegisterRoutes register routes
func RegisterRoutes(router *gin.Engine, service PostService) {
	c := controller{service}
	r := router.Group("/posts")
	r.Use(middleware.OptionalMiddleware())
	{
		r.GET("/", c.FetchProjectPosts)
		// r.GET("/:id", c.FetchCreatorPosts)
	}

	// r.Use(middleware.JwtMiddleware())
	// {
	// 	r.POST("/", c.CreateCreator)
	// }
}

func (con controller) FetchProjectPosts(c *gin.Context) {
	dto := dto.FetchProjectPostsDTO{}
	c.BindQuery(&dto)
	if dto.Limit <= 0 || dto.ProjectID <= 0 || dto.OrderBy == "" || dto.Page <= 0 || dto.Mature <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
	}
	fmt.Println(dto)
}
