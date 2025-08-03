package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jllovet/learning-go/09-tiago-backend-courses/000-ecom/config"
)

func CreateJWT(secret []byte, userID int) (string, error) {

	expirationTime := time.Second * time.Duration(config.InitializedConfig.JWTExpirationInSeconds)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.Itoa(userID),                  // subject -> the user id here
		"exi": time.Now().Add(expirationTime).Unix(), // expires in -> the number of seconds that the jwt will be valid
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
