package main

import(
	"github.com/gin-gonic/gin"
	"emp/database"
	"emp/controllers"
)

func loadDatabase() {
	database.LoadEnv()
	database.Connect()
	database.SyncDatabase()
}

func main() {
	loadDatabase()

	router := gin.Default()

	router.GET("/emp", controllers.GetUserDefualtMonth)

	// router.POST("/login", controllers.Login)

	// router.GET("/validate", controllers.Validate)

	router.Run(":3000")

}