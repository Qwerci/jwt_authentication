package models

import (
	"html"
	"jwt/database"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	//ID            *string   `gorm:"size:4; not null; unique" json:"id"`
	First_name    string    `gorm:"size: 25; not null" json:"first_name" validate:"required, min=5,max=25"`
	Last_name     string    `gorm:"size: 25; not null" json:"last_name" validate:"required, min=2, max=25"`
	Password      string    `gorm:"size:20; not null" json:"password" validate:"required, min=6, containany=!@#?*"`
	Email         string    `gorm:"size:100; not null; unique" json:"email" validate:"required, email"`
	// 
	User_type     string    `json:"user_type" validate:"required, eq=Director | eq=Operation_Manager | eq=HR | eq= Team_Leader | eq=QA | eq=System"`
	// Refresh_token *string   `json:"refresh_token"`
	CreatedAt     time.Time `json:"created_at"`
	UpdateAt      time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
	// User_id       string    `json:"user_id"`
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}


func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	return nil
}
