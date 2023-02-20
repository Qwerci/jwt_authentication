package main

import (
	"jwt_v2/controllers"
	"jwt_v2/database"

	"github.com/gin-gonic/gin"
)

func loadDatabase() {
	database.LoadEnv()
	database.Connect()
	database.SyncDatabase()
}

func main() {
	loadDatabase()

	router := gin.Default()

	router.POST("/signup", controllers.Signup)

	router.POST("/login", controllers.Login)

	router.GET("/validate", controllers.Validate)

	router.Run(":3000")

}
