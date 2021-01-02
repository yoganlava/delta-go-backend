package files

import (
	"main/internal/utility"

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
	var err error
	defer utility.ErrorHandleHTTP(c, err)
	file, err := c.FormFile("file")

	if err != nil {
		return
	}

	//Do stuff with file
	println(file)

}
