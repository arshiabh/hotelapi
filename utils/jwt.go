package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const Secret_key = "supersecret"

func VerifyToken(tokenstr string) (string, string, error) {
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token is invalid or expired")
		}
		return []byte(Secret_key), nil
	})
	if err != nil {
		return "", "", fmt.Errorf("token is invalid or expired")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", "", fmt.Errorf("token is invalid or expired")
	}
	//parse valid token and get this item we set then sent it to auth
	userID := claims["userID"].(string)
	email := claims["email"].(string)
	expired := claims["exp"].(float64)
	if time.Now().After(time.Unix(int64(expired), 0)) {
		return "", "", fmt.Errorf("token is expired")
	}
	return userID, email, nil
}

func GenarateToken(userID string, email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"email":  email,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	tokenstr, err := token.SignedString([]byte(Secret_key))
	if err != nil {
		log.Fatal("failed to created token")
	}
	return tokenstr
}
