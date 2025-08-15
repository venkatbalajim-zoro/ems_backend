package middleware

import (
	"log"
	"strings"
)

func Verify(text string) bool {
	if text == "" || !strings.HasPrefix(text, "Bearer ") {
		log.Println("Authorization code is required to proceed.")
		return false
	}

	token := strings.TrimPrefix(text, "Bearer ")
	if token == "" {
		log.Println("No authorization code is found.")
		return false
	}

	return true
}
