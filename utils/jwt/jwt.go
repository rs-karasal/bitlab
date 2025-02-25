package jwt

import (
	"my_super_project/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(cfg *config.Config, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_name": username,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(cfg.SecretKey))
}
