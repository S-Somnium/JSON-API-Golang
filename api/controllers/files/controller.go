package files

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	routes := router.Group("/files")
	{
		routes.POST("/", createFile)
		routes.GET("/:id", getFiles)
		routes.DELETE("/:id", deleteFiles)
	}
}
