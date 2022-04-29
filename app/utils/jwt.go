package utils

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(id int) (token string) {
	claims := jwt.StandardClaims{
		Id:        strconv.Itoa(int(id)),
		Issuer:    strconv.Itoa(int(id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("secret"))
	if err != nil {
		return
	}
	return token
}
