package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	routes := router.Group("/user")
	{
		routes.PATCH("/", updateUser)
		routes.DELETE("/:id", deleteUser)
		routes.GET("/", getAllUsers)
		routes.GET("/:id", getUser)
	}
}
