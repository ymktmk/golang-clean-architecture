package infrastructure

import (
	"github.com/ymktmk/golang-clean-architecture/interfaces/controllers"
	"github.com/labstack/echo"
)

func Routing() *echo.Echo {
	userController := controllers.NewUserController(NewSqlHandler())
	todoController := controllers.NewTodoController(NewSqlHandler())
	echo := echo.New()
	// routing
	echo.POST("/users/create", userController.Create)
	echo.POST("/todo",todoController.Create)
	return echo
}