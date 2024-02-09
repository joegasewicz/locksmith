package utilities

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type TokenClaims struct {
	jwt.RegisteredClaims
}

func CreateNewJWT(UserID uint) (string, error) {
	//claims := TokenClaims{
	//jwt.RegisteredClaims{
	//	Issuer:    "noto.com",
	//	Subject:   "authentication",
	//	Audience:  nil,
	//	ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * (24 * time.Hour))),
	//	NotBefore: jwt.NewNumericDate(time.Now()),
	//	IssuedAt:  jwt.NewNumericDate(time.Now()),
	//},
	//}
	claims := jwt.MapClaims{
		"user_id": UserID,
		"iss":     "noto.com",
		"sub":     "authentication",
		"exp":     jwt.NewNumericDate(time.Now().Add(7 * (24 * time.Hour))),
		"nbf":     jwt.NewNumericDate(time.Now()),
		"iat":     jwt.NewNumericDate(time.Now()),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(GetTokenSecret()))
	return tokenStr, err
}

func ParseToken(t string) (map[string]interface{}, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetTokenSecret()), nil
	})
	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			return claims, nil
		} else {
			return nil, errors.New("not a valid token")
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		return nil, errors.New("not a valid token")
	} else if ve.Errors&jwt.ValidationErrorMalformed != 0 {
		// expired
		return nil, errors.New("token expired")
	} else {
		return nil, errors.New("unable to parse token")
	}
}
