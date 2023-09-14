package utils

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func ComparePasswords(hashedPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	resp := map[string]string{
		"error": message,
	}
	_ = json.NewEncoder(w).Encode(resp)
}
