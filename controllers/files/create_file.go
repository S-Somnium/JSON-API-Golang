package files

import (
	"JSON-API-GOLANG/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createFile(c *gin.Context) {
	file := models.Files{}

	if err := c.ShouldBindJSON(&file); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(file)
	err := file.Create()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}
