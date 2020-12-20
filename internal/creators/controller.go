package creators

import (
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
	r := router.Group("/creator")
	r.POST("/:id", c.FetchCreator)
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
