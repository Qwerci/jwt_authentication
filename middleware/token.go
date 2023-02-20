package middleware

import(
	"github.com/golang-jwt/jwt/v4"
	"jwt_v2/models"
	"os"
	"time"
)

func GenerateJWT(user models.User) (string, error) {
	
	signMethod := jwt.SigningMethodHS256
	privateKey:= []byte(os.Getenv("JWT_PRIVATE_KEY"))

	
	claims := jwt.MapClaims{
		"sub":user.ID,
		"role": user.Role,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(signMethod, claims)

	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}