package controllers

import (
	"Golang-CleanArchitecture/domain"
	"Golang-CleanArchitecture/interfaces/database"
	"Golang-CleanArchitecture/usecase"
	"net/http"
	"strconv"
	"github.com/labstack/echo"
)

type UserController struct {
    Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
    return &UserController{
        Interactor: usecase.UserInteractor{
            UserRepository: &database.UserRepository{
                SqlHandler: sqlHandler,
            },
        },
    }
}

func (controller *UserController) CreateUser(c echo.Context) error {
	u := new(domain.User)
	if err := c.Bind(&u); err != nil {
		return err
	}

	user, err := controller.Interactor.Add(*u)
    if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) GetAllUsers(c echo.Context) error {
	users, err := controller.Interactor.Users()
    if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

func (controller *UserController) GetUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := controller.Interactor.Show(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}


