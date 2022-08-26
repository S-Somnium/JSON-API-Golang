package user

import (
	"JSON-API-GOLANG/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := models.LoginCheck(user.Email, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}
	c.SetCookie("token", token, 60*60*24, "/", "http://localhost:8081", true, true)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
