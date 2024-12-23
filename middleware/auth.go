package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const secret_key = "supersecret"

func JWTAuthentication(c *fiber.Ctx) error {
	token := c.GetRespHeader("X-Api-Token")
	if err := ParseJWTtoken(token); err != nil {
		return err
	}
	return nil
}

func ParseJWTtoken(tokenstr string) error {
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Printf("Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("Unauthorized")
		}
		return []byte(secret_key), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	}
	return fmt.Errorf("Unauthorized")
}


func GenarateToken(userid primitive.ObjectID, email string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uID":userid,
		"exp":time.Now().Add(time.Hour * 2).Unix(),
	},
	)
	return token.SignedString([]byte(secret_key))
}