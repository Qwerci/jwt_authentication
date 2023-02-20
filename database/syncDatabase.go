package database

import(
	"jwt_v2/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}