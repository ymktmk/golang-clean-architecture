package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/ymktmk/golang-clean-architecture/app/interfaces/controllers"

	"gopkg.in/go-playground/validator.v9"
)

func NewRouter() *echo.Echo {
	userController := controllers.NewUserController(NewSqlHandler())
	todoController := controllers.NewTodoController(NewSqlHandler())
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.POST},
	}))
	e.Validator = &CustomValidator{Validator: validator.New()}
	// register & login & logout
	e.POST("/api/register", userController.Register, issueCookie)
	e.POST("/api/login", userController.Login, issueCookie)
	e.GET("/api/logout", userController.Logout)
	// middleware
	g := e.Group("/api", VerifyToken)
	// user
	g.GET("/user", userController.Show)
	g.PATCH("/user", userController.Update)
	// todos
	g.POST("/todos", todoController.Create)
	g.GET("/todos", todoController.All)
	g.GET("/todos/:id", todoController.Show)
	g.PATCH("/todos/:id", todoController.Update)
	return e
}
