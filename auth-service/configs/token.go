package configs

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var secret = GetEnv("JWT_SECRET", "sample")

func GenerateToken(username string, id int) (string, error) {
	claims := jwt.MapClaims{
		"username":    username,
		"employee_id": id,
		"expiry":      time.Now().Add(time.Hour * 24).Unix(),
	}

	method := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err := method.SignedString([]byte(secret))
	return result, err
}
