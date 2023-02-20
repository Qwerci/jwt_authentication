package controllers

import (
	"golang.org/x/crypto/bcrypt"
	"jwt_v2/database"
	"jwt_v2/models"
	"jwt_v2/middleware"
	"net/http"
	"github.com/gin-gonic/gin"
)


// var jwtKey = []byte("my_secret_key")
// var refreshJwtKey = []byte("my_refresh_secret_key")

func Signup(c *gin.Context) {
	// Get the inputs needed for registration

	var body models.User
	if err:= c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read the request body."})
		return
	}

	// Validate the user's email and password

	if body.Email == "" || body.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}
	

	// Hash the user's password before storing it in the database

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error hashing password"})
		return
	}
	body.Password = string(hashedPassword)

	// Create a new user in the database
	user := models.User{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, Password: string(hashedPassword), Role: body.Role}
	result := database.DB.Create(&user)

	if result.Error !=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Failed to create user"})
	}

	// Save the user in the database

	// err = database.DB.Save(&user).Error
	// if err!= nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Error saving user"})
	// 	return
	// }

	// Return a success message
	 c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})

	
}




func Login(c *gin.Context) {
	var user models.User
	var body models.LoginBody
	if err:= c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read the request body."})
		return
	}

	// Find the user in the database

	database.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or Password"})
		return
	}
	// Compare the provided password with the hashed password in the database

	err  := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err!= nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or Password"})
		return
	}

	// Generate a JWT token and set it in an HttpOnly cookie
	
	token, err := middleware.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	// refreshToken, err := generateRefreshToken(user.ID)
    // if err != nil {
    //     c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating refresh token"})
    //     return
	// }

	cookie := http.Cookie{
		Name: "jwt",
		Value: token,
		MaxAge: 3600 * 24 * 30,
		HttpOnly: true,
	}
	
	// refreshCookie := http.Cookie{
    //     Name:     "refreshToken",
    //     Value:    refreshToken,
    //     HttpOnly: true,
	// }

	http.SetCookie(c.Writer,&cookie)
	// http.SetCookie(c.Writer, &refreshCookie)

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Login succeful"})
}

func Validate(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"message": "I'm in"})
}
// func authMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Get access token from HTTP-only cookie
// 		token, err := c.Cookie("access_token")
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H)
// 		}
// 	}
// } 