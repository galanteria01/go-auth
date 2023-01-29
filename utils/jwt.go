package utils

import (
	"example/go-auth/configs"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = configs.EnvSecretKey()

func GenerateJWT(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodEdDSA)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(10 * time.Hour)
	claims["authorized"] = true
	claims["user"] = username
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "signing error", err
	}

	return tokenString, nil
}
