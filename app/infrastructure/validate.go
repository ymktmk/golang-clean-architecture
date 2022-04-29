package infrastructure

import (
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	Validator *validator.Validate
}

// c.Validateで使えるようになる
func (cv *CustomValidator) Validate(i interface{}) error {
	// cv.Validator.RegisterValidation("email_check", Email)
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

// func Email(fl validator.FieldLevel) bool {
// 	return fl.Field().String()
// }
