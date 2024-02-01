package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

const secretKey = "secretsauce"

func GenerateJWT(email string, userId int64) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	}).SignedString([]byte(secretKey))
}

func ValidateToken(token string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, errors.New("Could not parse token")
	}

	if !parsedToken.Valid {
		return nil, errors.New("Invalid token")
	}

	return parsedToken, nil
}

func GetUserIdFromToken(jwtToken *jwt.Token) (int64, error) {
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return -1, errors.New("Unable to retrieve claims")
	}

	userId := int64(claims["userId"].(float64))
	return userId, nil
}
