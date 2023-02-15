package controller

import(
	"jwt/models"
	"jwt/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"jwt/helper"

)

// type value struct{
// 	models.User
// }

func Register( context *gin.Context) {
	var input models.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err!= nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	user := models.User{
		First_name: input.First_name,
		Last_name: input.Last_name,
		Password: input.Password,
		Email: input.Email,
		User_type: input.User_type,
	}

	savedUser, err := user.Save()

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})

}


func ValidatePassword(user models.User, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))	
}

func FindUserByEmail(email string) (models.User, error) {
	var user models.User
	err := database.Database.Where("email =?", email).Find(&user).Error
	if err!= nil {
		return models.User{}, err
	}
	return user, nil
}


func Login(context *gin.Context){
	var input models.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err!= nil {
		return
	}

	user, err := FindUserByEmail(input.Email)

	if err!= nil{
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ValidatePassword(user,input.Password)

	if err!= nil{
		context.JSON(http.StatusBadRequest, gin.H{"error":  "Invalid login credentials"})
		return
	}

	jwt, err := helper.GenerateJWT(user)

	if err!= nil{
		context.JSON(http.StatusBadRequest, gin.H{"error":"Error generating token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
