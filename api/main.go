package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	router.Run(":8080")
}

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}
}
