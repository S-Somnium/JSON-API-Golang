package files

import (
	"JSON-API-GOLANG/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func deleteFiles(c *gin.Context) {
	id := c.Param("id")
	userID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err = models.DeleteAllFiles(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user files deleted!"})
}
