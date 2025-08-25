package utils

import (
	"auth-service/configs"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(username string, id int) (string, error) {
	var secret = configs.GetEnv("JWT_SECRET", "sample")
	
	claims := jwt.MapClaims{
		"username":    username,
		"employee_id": id,
		"expiry":      time.Now().Add(time.Hour * 24).Unix(),
	}

	method := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err := method.SignedString([]byte(secret))
	return result, err
}
