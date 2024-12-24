package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const secret_key = "supersecret"

func JWTAuthentication(c *fiber.Ctx) error {
	token := c.GetReqHeaders()["X-Api-Token"]
	tokenstr := token[0]
	if tokenstr == ""{
		return c.JSON(fiber.Map{"error":"unauthorized"})
	}
	email, err := ParseJWTtoken(tokenstr)
	if  err != nil {
		return err
	}
	c.Set("email", email)
	return c.Next()
}

func ParseJWTtoken(tokenstr string) (string, error) {
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Printf("Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("Unauthorized")
		}
		return []byte(secret_key), nil
	})
	if err != nil {
		return "",err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", err
	} 
	email := claims["email"].(string)
	return email, nil	
}


func GenarateToken(email string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":email,
		"exp":time.Now().Add(time.Hour * 2).Unix(),
	},
	)
	return token.SignedString([]byte(secret_key))
}