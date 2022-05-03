package utils_test

import (
	"fmt"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/ymktmk/golang-clean-architecture/app/utils"
)

type Claims struct {
	jwt.StandardClaims
}

func TestCreateToken(t *testing.T) {
	// id=1のjwtTokenを作成
	tokenString := utils.CreateToken(1)
	fmt.Println(tokenString)
	// jwtTokenを検証
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return
	}
	claims := token.Claims.(*Claims)
	id := claims.Issuer
	assert.Equal(t, "1", id)
}
