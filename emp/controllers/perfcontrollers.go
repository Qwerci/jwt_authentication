package controllers

import (
	"emp/database"
	"emp/models"
	"net/http"
	"github.com/gin-gonic/gin"
)


func GetUserDefualtMonth(c *gin.Context){
	// Declare a variable use the latest month and year in the database as defaults
	var latest models.Performance
	database.DB.Order("month Desc, year Desc").First(&latest)
	month := c.DefaultQuery("month", latest.Month)
	year := c.DefaultQuery("year", latest.Year)

	// Query the database for all users with the specified month and year
	
	var body []models.Performance
	database.DB.Select("id,first_name,middle_name,last_name,quality,competency,rtr,attendance,project_target_score,total,month,year").Where("month=? AND year=?", month, year).Find(&body)

	 // Return the users as JSON
	 c.JSON(http.StatusOK, gin.H{"post": body,})

}
