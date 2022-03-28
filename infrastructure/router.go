package infrastructure

import (
	"Golang-CleanArchitecture/interfaces/controllers"
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