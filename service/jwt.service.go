package service

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

const USER_ID_SCHEMA = "user_id"

func GenerateToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})

	tokenString, err := token.SignedString([]byte("idev"))
	if err != nil {
		panic(err)
	}
	return tokenString
}

func ValidateToken(tokenString string) (isValid bool, err error) {
	token, _err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, _isValid := token.Method.(*jwt.SigningMethodHMAC); !_isValid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])
		}
		return []byte("idev"), nil
	})
	isValid = token != nil && token.Valid
	err = _err
	return
}
