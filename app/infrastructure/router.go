package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/ymktmk/golang-clean-architecture/app/interfaces/controllers"
	"github.com/ymktmk/golang-clean-architecture/app/interfaces/database"
	"github.com/ymktmk/golang-clean-architecture/app/usecase/interactor"

	"gopkg.in/go-playground/validator.v9"
)

func NewRouter() *echo.Echo {
	userController := controllers.NewUserController(NewSqlHandler())
	todoController := controllers.NewTodoController(interactor.NewTodoInteractor(database.NewTodoRepository(NewSqlHandler())))
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Validator = &CustomValidator{Validator: validator.New()}
	// register & login & logout
	e.POST("/api/register", userController.Register, issueCookie)
	e.POST("/api/login", userController.Login, issueCookie)
	e.GET("/api/logout", userController.Logout)
	// middleware
	g := e.Group("/api", VerifyToken)
	// user
	g.GET("/users", userController.Show)
	g.PATCH("/users", userController.Update)
	// todos
	g.POST("/todos", todoController.Create)
	g.GET("/todos", todoController.All)
	g.GET("/todos/:id", todoController.Show)
	g.PATCH("/todos/:id", todoController.Update)
	return e
}
