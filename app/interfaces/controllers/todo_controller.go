package controllers

import (
	"net/http"
	"strconv"

	"github.com/ymktmk/golang-clean-architecture/app/domain"
	"github.com/ymktmk/golang-clean-architecture/app/interfaces/database"
	"github.com/ymktmk/golang-clean-architecture/app/usecase"

	"github.com/labstack/echo"
)

type TodoController struct {
	Interactor usecase.TodoInteractor
}

func NewTodoController(sqlHandler database.SqlHandler) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: &database.TodoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// todo作成
func (controller *TodoController) Create(c echo.Context) (err error) {
	tcr := new(domain.TodoCreateRequest)
	if err = c.Bind(tcr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(tcr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userIdString := c.Get("id").(string)
	userId, _ := strconv.Atoi(userIdString)
	// DTO
	t := &domain.Todo{
		Name: tcr.Name,
		UserID: userId,
	}
	todo, err := controller.Interactor.Create(t)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

func (controller *TodoController) Update(c echo.Context) (err error) {
	tcr := new(domain.TodoCreateRequest)
	if err = c.Bind(tcr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(tcr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	// DTO
	t := &domain.Todo{Name: tcr.Name}
	todo, err := controller.Interactor.Update(id,t)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

// 1つのtodo取得
func (controller *TodoController) Show(c echo.Context) (err error) {
	idString := c.Param("id")
	id, _ := strconv.Atoi(idString)
	todo, err := controller.Interactor.TodoById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

// 全てのtodo取得
func (controller *TodoController) All(c echo.Context) (err error) {
	userIdString := c.Get("id").(string)
	userId, _ := strconv.Atoi(userIdString)
	todos, err := controller.Interactor.TodosById(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, todos)
}