package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret []byte = []byte("secret")

func CreateToken(payload jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString(secret)
}

func CreateExpiredToken(payload jwt.MapClaims, exp int64) (string, error) {
	payload["exp"] = exp
	return CreateToken(payload)
}

func DecodeJWT(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
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

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.New("exp claim is missing or invalid")
	}

	if int64(exp) < time.Now().Unix() {
		return nil, errors.New("token has expired")
	}

	return claims, nil
}
