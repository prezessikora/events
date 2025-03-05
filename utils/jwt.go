package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secretKey"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 6).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		fmt.Println("error parsing token")
		return 0, err
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		fmt.Println("token invalid")
		return 0, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalud token claims")
	}
	fmt.Println(parsedToken.Claims)
	// email, _ := claims["email"].(string)
	userId, ok := claims["userId"].(float64)

	return int(userId), nil
}
