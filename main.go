package main

import (
	"JSON-API-GOLANG/controllers/files"
	"JSON-API-GOLANG/controllers/user"
	"JSON-API-GOLANG/middlewares"
	"os"

	log "github.com/sirupsen/logrus"

	"JSON-API-GOLANG/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.New()
	router.Use(middlewares.Logger())
	router.Use(middlewares.CORSMiddleware())
	router.POST("/login", user.Login)
	router.POST("/register", user.CreateUser)
	router.Use(middlewares.IsAuthenticated())
	files.RegisterRoutes(router)
	user.RegisterRoutes(router)
	router.Run(":8080")
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}
	models.ConnectDataBase()
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}
