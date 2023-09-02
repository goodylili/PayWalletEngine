package http

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"os"
	"time"
)

// validateJWT - validates an incoming jwt token
func validateJWT(accessToken string) bool {
	err := godotenv.Load()
	if err != nil {
		return false
	}

	var mySigningKey = []byte(os.Getenv("JWT_KEY"))

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not validate auth token")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return false
	}

	return token.Valid
}

func generateJWT() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	secretKey := []byte(os.Getenv("JWT_KEY"))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(15 * time.Minute)
	claims["authorized"] = true
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
