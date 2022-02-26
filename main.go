package main

import (
	"Golang-CleanArchitecture/infrastructure"
	"github.com/labstack/echo/middleware"
)

func main() {
	echo := infrastructure.Routing()
	echo.Use(middleware.CORS())
	echo.Logger.Fatal(echo.Start(":8080"))
}