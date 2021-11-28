package utils

import (
	"time"
	"user_service/models"

	"github.com/golang-jwt/jwt"
)

func TokenJwt(user models.RegisterUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(48 * time.Hour).Unix(),
	})

	strToken, err := token.SignedString([]byte("secret-key"))

	if err != nil {
		return "", err
	}

	return strToken, nil

}
