package tiers

import (
	"main/internal/utility"
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
	var err error
	defer utility.ErrorHandleHTTP(c, err)
	id, _ := strconv.Atoi(c.Param("id"))
	tiers, err := con.service.FetchProjectTiers(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, tiers)
}

func (con controller) fetchTier(c *gin.Context) {
	var err error
	defer utility.ErrorHandleHTTP(c, err)
	id, _ := strconv.Atoi(c.Param("id"))
	tiers, err := con.service.FetchTier(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, tiers)
}
