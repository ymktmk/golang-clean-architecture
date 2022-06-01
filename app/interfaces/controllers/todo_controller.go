package controllers

import (
	"net/http"
	"strconv"

	"github.com/ymktmk/golang-clean-architecture/app/domain"
	"github.com/ymktmk/golang-clean-architecture/app/usecase/input"

	"github.com/labstack/echo"
)

type TodoController struct {
	Usecase input.TodoUsecase
}

func NewTodoController(usecase input.TodoUsecase) *TodoController {
	return &TodoController{
		Usecase: usecase,
	}
}

func (controller *TodoController) Create(c echo.Context) (err error) {
	tcr := new(domain.TodoCreateRequest)
	if err = c.Bind(tcr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(tcr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userId, _ := strconv.ParseUint(c.Get("id").(string), 10, 32)
	inputData := input.CreateTodo{
		Name:   tcr.Name,
		UserID: uint(userId),
	}
	todo, err := controller.Usecase.Create(inputData)
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

	todoId, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	userId, _ := strconv.ParseUint(c.Get("id").(string), 10, 32)
	// todoIdからtodoを取得する
	td, err := controller.Usecase.GetTodo(input.GetTodo{
		TodoID: uint(todoId),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// userIDが一致するかどうか
	if uint(userId) != td.UserID {
		return echo.NewHTTPError(http.StatusBadRequest, "不当なユーザーのリクエストです")
	}
	todo, err := controller.Usecase.Update(input.UpdateTodo{
		ID:   uint(todoId),
		Name: tcr.Name,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

// 1つのtodo取得
func (controller *TodoController) Show(c echo.Context) (err error) {
	todoId, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	todo, err := controller.Usecase.GetTodo(input.GetTodo{TodoID: uint(todoId)})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

// 全てのtodo取得
func (controller *TodoController) All(c echo.Context) (err error) {
	userId, _ := strconv.ParseUint(c.Get("id").(string), 10, 32)
	todos, err := controller.Usecase.GetAllTodos(input.GetAllTodos{UserID: uint(userId)})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, todos)
}
