package utils

import (
	"OneTix/configs"
	"OneTix/models"
	"OneTix/structs"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(user models.MstUser) (tokenString string, err error) {
	env := configs.LoadEnv()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, structs.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{},
		ID:               user.ID,
		Role:             user.Role,
	})

	tokenString, err = token.SignedString([]byte(env.JWTSecret))

	return
}

func VerifyToken(tokenString string) (*structs.JWTClaims, error) {
	env := configs.LoadEnv()

	token, err := jwt.ParseWithClaims(tokenString, &structs.JWTClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(env.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*structs.JWTClaims); ok && token.Valid {
		if claims.ExpiresAt != nil && claims.ExpiresAt.Time.Before(time.Now()) {
			return nil, fmt.Errorf("token has expired")
		}
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
