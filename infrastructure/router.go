package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/ymktmk/golang-clean-architecture/interfaces/controllers"
	"gopkg.in/go-playground/validator.v9"
)

func Routing() *echo.Echo {
	userController := controllers.NewUserController(NewSqlHandler())
	todoController := controllers.NewTodoController(NewSqlHandler())
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	// routing
	e.POST("/users/create", userController.Create)
	e.POST("/todo",todoController.Create)
	return e
}