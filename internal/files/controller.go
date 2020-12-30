package files

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	service FileService
}

func RegisterRoutes(router *gin.Engine, service FileService) {
	// c := controller{service}
	// u := router.Group("/files")
}

func (con controller) uploadFile(c *gin.Context) {
	file, err := c.FormFile("file")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No file is received",
		})
		return
	}

	//Do stuff with file
	println(file)

}
