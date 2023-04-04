package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(validFor time.Duration, userID string, secretKey string) (string, error) {
	// create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// set the claims for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID
	claims["exp"] = time.Now().Add(validFor).Unix()

	// sign the token with the secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string, secretKey string) (string, error) {
	// parse the token string
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// return the secret key
		return []byte(secretKey), nil
	})

	// check for parsing errors
	if err != nil {
		return "", err
	}

	// check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check if the token has expired
		exp := time.Unix(int64(claims["exp"].(float64)), 0)
		if exp.Before(time.Now()) {
			return "", errors.New("token has expired")
		}

		// return the user ID
		return claims["id"].(string), nil
	} else {
		return "", errors.New("invalid token")
	}
}