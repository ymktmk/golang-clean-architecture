package main

import (
	"github.com/labstack/echo/middleware"

	"github.com/ymktmk/golang-clean-architecture/infrastructure"
)

func main() {
	e := infrastructure.NewRouter()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.Fatal(e.Start(":9000"))
}
