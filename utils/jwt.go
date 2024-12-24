package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const Secret_key = "supersecret"

func VerifyToken(tokenstr string) (string, error) {
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token is invalid or expired")
		}
		return []byte(Secret_key), nil
	})
	if err != nil {
		return "", fmt.Errorf("token is invalid or expired")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("token is invalid or expired")
	}
	email := claims["email"].(string)
	expired := claims["exp"].(float64)
	if time.Now().After(time.Unix(int64(expired),0)) {
		return "", fmt.Errorf("token is expired")
	}
	return email, nil
}

func GenarateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(Secret_key))
}
