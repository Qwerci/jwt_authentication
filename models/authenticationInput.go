package models

type AuthenticationInput struct {
    First_name string `json:"first_name" binding:"required"`
	Last_name string `json:"last_name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email string `json:"email" binding:"required"`
	User_type string `json:"user_type" binding:"required"`
    
}