package jwt

import (
	"AILN/app/common"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func SignToken(userID uint) (string, error) {
	secretKey := common.CONFIG.String("jwt.secretKey")
	iat := common.CONFIG.Int("jwt.Issuer")
	seconds := common.CONFIG.Int("ExpireSeconds")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    iat + seconds,
		"iat":    iat,
		"userID": userID,
	})

	return token.SignedString([]byte(secretKey))
}

func Parse(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(common.CONFIG.String("jwt.secretKey")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("parse token: invalid claim type")
	}

	if err = claims.Valid(); err != nil {
		return nil, fmt.Errorf("invalid claims: %v", err)
	}

	return claims, nil
}
