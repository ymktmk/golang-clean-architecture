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
		// cookie, _ := c.Cookie("jwt")

		// token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 	return []byte("secret"), nil
		// })

		// if err != nil || !token.Valid {
		// 	return err
		// }

		// claims := token.Claims.(*Claims)
		// id := claims.Issuer

		// var user domain.User
		// db := NewSqlHandler()
		// if err = db.Where("id = ?", id).First(&user).Error; err != nil {
		// 	return err
		// }
		// c.Set("user", &user)

		// if err := next(c); err != nil {
		// 	c.Error(err)
		// }
		return nil
	}
}
