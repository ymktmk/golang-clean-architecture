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
	e.POST("/api/register", userController.Register)
	e.POST("/api/login", userController.Login)
	e.GET("/api/logout", userController.Logout)
	g := e.Group("/api", verifyToken)
	g.GET("/user", userController.Show)
	e.POST("/todo", todoController.Create)
	return e
}
