package api

import (
	"github.com/golang-jwt/jwt"
	"github.com/ssoql/auth-service/config"
	"log"
)

func init() {
	if len(config.SecurityKey) == 0 {
		log.Fatal("please set SecurityKey environment value")
	}
}

func IsTokenValid(token string) bool {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecurityKey), nil
	})
	if err != nil || !parsedToken.Valid {
		return false
	}

	return true
}
