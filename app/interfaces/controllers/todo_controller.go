package controllers

import (
	"net/http"

	"github.com/ymktmk/golang-clean-architecture/app/interfaces/database"
	"github.com/ymktmk/golang-clean-architecture/app/usecase"

	"github.com/labstack/echo"
)

type TodoController struct {
	Interactor usecase.TodoInteractor
}

// 依存性を定義する
func NewTodoController(sqlHandler database.SqlHandler) *TodoController {
	return &TodoController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: &database.TodoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

// シフトを作成する
func (controller *TodoController) Create(c echo.Context) (err error) {
	// uidを取得
	uid := c.Get("uid").(string)
	// userを取得する
	user, err := controller.Interactor.UserByUid(uid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}