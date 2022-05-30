package infrastructure

import (
	// "net/http"
	// "time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	// "github.com/ymktmk/golang-clean-architecture/app/utils"
)

type Claims struct {
	jwt.StandardClaims
}

func issueCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// 前処理
		if err := next(c); err != nil {
			c.Error(err)
		}
		// 後処理
		// userId, _ := c.Get("userId").(int)
		// token := utils.CreateToken(userId)
		// cookie := http.Cookie{
		// 	Name:     "jwt",
		// 	Value:    token,
		// 	Path:     "/",
		// 	Expires:  time.Now().Add(time.Hour * 24),
		// 	HttpOnly: true,
		// }
		// c.SetCookie(&cookie)
		return nil
	}
}

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
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
