package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

var signingKey = []byte(viper.GetString("jwt.signingKey"))

type JwtCustClaims struct {
	Name string
	jwt.RegisteredClaims
}

func GenerateToken(name string) (string, error) {
	iJwtCustClaims := JwtCustClaims{
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("jwt.tokenExpire") * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "Token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustClaims)
	return token.SignedString(signingKey)
}

func ParseToken(token string) (JwtCustClaims, error) {
	iJwtCustClaims := JwtCustClaims{}
	t, err := jwt.ParseWithClaims(token, &iJwtCustClaims, func(t *jwt.Token) (any, error) {
		return signingKey, nil
	})
	if err == nil && !t.Valid {
		err = errors.New("Invalid Token")
	}
	return iJwtCustClaims, err
}

func IsTokenValid(token string) bool {
	_, err := ParseToken(token)
	return err == nil
}
