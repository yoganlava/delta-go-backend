package projects

import (
	"main/internal/dto"
	"main/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service *ProjectService
}

func RegisterRoutes(router *gin.Engine, service *ProjectService) {
	c := controller{service}
	r := router.Group("/projects")
	r.Use(middleware.OptionalMiddleware())
	{
		r.GET("/:url", c.FetchProject)
	}
	r.GET("/:url/isAvailable", c.FetchIsURLAvailable)
	r.Use(middleware.JwtMiddleware())
	{
		r.POST("/", c.CreateProject)

	}
}

func (con controller) FetchProject(c *gin.Context) {
	// Will change later to accomodate for projects with custom url
	// id, _ := strconv.Atoi()
	project, err := con.service.FetchProject(c.Param("url"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "プロジェクトが見つかりません",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"project": project,
	})
}

func (con controller) FetchIsURLAvailable(c *gin.Context) {
	available, err := con.service.isPageURLAvailable(c.Param("url"))
	if err != nil {
		c.JSON(http.StatusFound, gin.H{
			"available": false,
			"reason":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusContinue, gin.H{
		"available": available,
	})
	return

}

func (con controller) CreateProject(c *gin.Context) {
	var project dto.CreateProjectDTO
	if err := c.BindJSON(&project); err != nil {
		c.AbortWithStatus(500)
		return
	}
	userID, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Creator
	creatorID, err := con.service.GetCreatorIDFromUserID(userID.(int))
	if err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}
	project.CreatorID = creatorID
	err = con.service.CreateProject(project)
	if err != nil {
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Success",
	})
}
