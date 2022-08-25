package files

import (
	"JSON-API-GOLANG/api/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getFiles(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	files, err := models.GetAllFiles(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, *files)

}
