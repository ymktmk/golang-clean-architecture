package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/ymktmk/golang-clean-architecture/app/interfaces/controllers"

	"gopkg.in/go-playground/validator.v9"
)

func NewRouter() *echo.Echo {
	userController := controllers.NewUserController(NewSqlHandler())
	todoController := controllers.NewTodoController(NewSqlHandler())
	e := echo.New()
	e.Validator = &CustomValidator{Validator: validator.New()}
	// routing
	e.POST("/register", userController.Register)
	// e.POST("/login", userController.Login)
	// e.POST("/logout", userController.Logout)
	e.GET("/user", userController.Show)

	e.POST("/todo", todoController.Create)
	return e
}
