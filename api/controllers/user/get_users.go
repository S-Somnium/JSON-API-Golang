package user

import (
	"JSON-API-GOLANG/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, *users)
}
