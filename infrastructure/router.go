package infrastructure

import (
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"

	"github.com/ymktmk/golang-clean-architecture/interfaces/controllers"
)

func NewRouter() *echo.Echo {
	userController := controllers.NewUserController(NewSqlHandler())
	todoController := controllers.NewTodoController(NewSqlHandler())
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	// routing
	e.POST("/users/create", userController.Create)
	e.GET("/user", userController.Show)

	e.POST("/todo", todoController.Create)
	return e
}
