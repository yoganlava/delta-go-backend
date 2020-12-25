package tiers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service TierService
}

// RegisterRoutes register routes
func RegisterRoutes(router *gin.Engine, service TierService) {
	c := controller{service}
	u := router.Group("/tiers")
	// Subject to change
	u.GET("/project/:id", c.fetchProjectTiers)
	u.GET("/:id", c.fetchProjectTiers)

}

func (con controller) fetchProjectTiers(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tiers, err := con.service.FetchProjectTiers(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, tiers)
}

func (con controller) fetchTier(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tiers, err := con.service.FetchTier(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, tiers)
}
