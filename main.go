package main

import (
    "jwt/database"
    "jwt/models"
	"jwt/controller"
	"fmt"
	"github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
    "log"
)

func main() {
    loadEnv()
    loadDatabase()
	serveApplication()
}

func loadDatabase() {
    database.Connect()
    database.Database.AutoMigrate(&models.User{})


    // database.Database.AutoMigrate(&model.Entry{})
}

func loadEnv() {
    err := godotenv.Load(".env.local")
    if err != nil {
        log.Fatal("Error loading .env file")
    } 
}

func serveApplication() {
	router := gin.Default()

	publicRouter := router.Group("/auth")
	publicRouter.POST("/register", controller.Register)
	publicRouter.POST("/login", controller.Login)


	router.Run(":8000")
	fmt.Println("Server running on port 8000")

}