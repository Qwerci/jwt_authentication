package models

import(
	"gorm.io/gorm"
	
)


type User struct {
		gorm.Model
		FirstName string    `gorm:"not null" json:"first_name" binding:"required"`
		LastName  string    `gorm:"not null" json:"last_name" binding:"required"`
		Email     string    `gorm:"unique ;not null" json:"email" binding:"required"`
		Password  string    `gorm:"not null" json:"password" binding:"required,min=8"`
		Role      string    `gorm:"not null" json:"role" binding:"required"`
		Token     string    `json:"-"`
		
}

type LoginBody struct{
	Email    string  
	Password string  	

}
	







