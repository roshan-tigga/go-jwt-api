package helpers

import (
	"auth/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return token.SignedString(config.JWT_KEY)
}
