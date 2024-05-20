package utils

import (
	"errors"
	"ex01/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(payload jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString(config.Config.Secret)
}

func DecodeJWT(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return config.Config.Secret, nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	if exp, ok := claims["exp"].(int64); ok && exp < time.Now().Unix() {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}
