package infrastructure

import (
	"Golang-CleanArchitecture/interfaces/controllers"
	"github.com/labstack/echo"
)

func Routing() *echo.Echo {
	userController := controllers.NewUserController(NewSqlHandler())
	echo := echo.New()
	// routing
	echo.POST("/users/create", userController.CreateUser)
	echo.GET("/users/:id", userController.GetUser)
	echo.GET("/users",userController.GetAllUsers)
	return echo
}