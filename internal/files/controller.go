package files

import (
	"github.com/gin-gonic/gin"
)

type controller struct {
	service FileService
}

func RegisterRoutes(router *gin.Engine, service FileService) {
	// c := controller{service}
	// u := router.Group("/files")
}
