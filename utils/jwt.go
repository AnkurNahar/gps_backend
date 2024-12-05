package utils

import (
	//"time"
	"errors"
	jwt "github.com/golang-jwt/jwt/v5"
	"log"
)

var secretKey = []byte("thisIsOurSecret!Shhhhh")

func GenerateJWT(userID string) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		//"exp":    time.Now().Add(time.Hour * 24).Unix(), //add expiration later
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		log.Println("Error generating token:", err)
		return "", err
	}
	//log.Println(signedToken)
	return signedToken, nil
}



func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}