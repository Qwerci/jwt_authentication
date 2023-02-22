package database

import(
	"emp/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.Performance{})
}