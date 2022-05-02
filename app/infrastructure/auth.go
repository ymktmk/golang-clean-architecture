package infrastructure

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type Claims struct {
	jwt.StandardClaims
}

func verifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, _ := c.Cookie("jwt")
		tokenString := cookie.Value
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil || !token.Valid {
			return err
		}
		claims := token.Claims.(*Claims)
		id := claims.Issuer
		c.Set("id", id)
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}
