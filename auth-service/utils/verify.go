package utils

import (
	"auth-service/configs"
	"errors"
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt"
)

func Verify(text string) (string, int, error) {
	var secret = configs.GetEnv("JWT_SECRET", "sample")

	if text == "" || !strings.HasPrefix(text, "Bearer ") {
		return "", 0, errors.New("authorization code is required")
	}

	tokenStr := strings.TrimPrefix(text, "Bearer ")
	if tokenStr == "" {
		return "", 0, fmt.Errorf("empty token")
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", 0, fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", 0, fmt.Errorf("invalid token")
	}

	if exp, ok := claims["expiry"].(float64); ok {
		if int64(exp) < jwt.TimeFunc().Unix() {
			return "", 0, fmt.Errorf("token expired")
		}
	}

	username, ok1 := claims["username"].(string)
	empIDFloat, ok2 := claims["employee_id"].(float64)

	if !ok1 || !ok2 {
		return "", 0, fmt.Errorf("username or employee_id not found in claims")
	}

	return username, int(empIDFloat), nil
}
