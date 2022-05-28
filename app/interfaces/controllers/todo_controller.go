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

// todo作成
func (controller *TodoController) Create(c echo.Context) (err error) {
	tcr := new(domain.TodoCreateRequest)
	if err = c.Bind(tcr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(tcr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	userId, _ := strconv.Atoi(c.Get("id").(string))
	// DTO
	inputData := input.CreateTodo{
		Name:   tcr.Name,
		UserID: userId,
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

	todoId, _ := strconv.Atoi(c.Param("id"))
	userId, _ := strconv.Atoi(c.Get("id").(string))
	// todoIdからtodoを取得する
	td, err := controller.Usecase.GetTodo(input.GetTodo{
		TodoID: todoId,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// userIDが一致するかどうか
	if userId != td.UserID {
		return echo.NewHTTPError(http.StatusBadRequest, "不当なユーザーのリクエストです")
	}
	todo, err := controller.Usecase.Update(input.UpdateTodo{
		ID:   todoId,
		Name: tcr.Name,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

// 1つのtodo取得
func (controller *TodoController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := controller.Usecase.GetTodo(input.GetTodo{TodoID: id})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, todo)
}

// 全てのtodo取得
func (controller *TodoController) All(c echo.Context) (err error) {
	userId, _ := strconv.Atoi(c.Get("id").(string))
	todos, err := controller.Usecase.GetAllTodos(input.GetAllTodos{UserID: userId})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, todos)
}
