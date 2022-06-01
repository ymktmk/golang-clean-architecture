package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"

	"github.com/ymktmk/golang-clean-architecture/app/domain"
	"github.com/ymktmk/golang-clean-architecture/app/usecase/input"
	"github.com/ymktmk/golang-clean-architecture/app/utils"
)

type UserController struct {
	Usecase input.UserUsecase
}

func NewUserController(usecase input.UserUsecase) *UserController {
	return &UserController{
		Usecase: usecase,
	}
}

func (controller *UserController) Register(c echo.Context) (err error) {
	ucr := new(domain.UserCreateRequest)
	if err = c.Bind(ucr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(ucr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// 同じEmailの人がいないか確認する && UIDも
	inputD := input.GetExistUserByEmail{
		Email: ucr.Email,
	}
	users, err := controller.Usecase.ExistUserByEmail(inputD)
	if len(users) != 0 {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// パスワードをハッシュ化
	password, _ := bcrypt.GenerateFromPassword([]byte(ucr.Password), 10)
	inputData := input.CreateUser{
		Name:     ucr.UserName,
		Email:    ucr.Email,
		Password: password,
	}
	// 同じメールアドレス、uidでerr返ってくる → 同じものを挿入したときidは進む
	user, err := controller.Usecase.Create(inputData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token := utils.CreateToken(int(user.ID))
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		Path:     "/",
		HttpOnly: true,
	}
	c.SetCookie(&cookie)
	return c.JSON(http.StatusOK, user)
}

// ログイン
func (controller *UserController) Login(c echo.Context) (err error) {
	ulr := new(domain.UserLoginRequest)
	if err = c.Bind(ulr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(ulr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	inputData := input.GetUserByEmail{
		Email: ulr.Email,
	}
	user, err := controller.Usecase.UserByEmail(inputData)
	if err != nil {
		return err
	}
	
	// これ必須
	// if err := bcrypt.CompareHashAndPassword(user.Password, []byte(ulr.Password)); err != nil {
	// 	return err
	// }
	
	token := utils.CreateToken(int(user.ID))
	
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	c.SetCookie(&cookie)
	return c.String(http.StatusOK, "success login !")
}


func (controller *UserController) Logout(c echo.Context) (err error) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}
	c.SetCookie(&cookie)
	return c.String(http.StatusOK, "success logout !")
}


func (controller *UserController) Show(c echo.Context) (err error) {
	id, _ := strconv.Atoi(c.Get("id").(string))
	inputData := input.GetUserById{
		ID: uint(id),
	}
	user, err := controller.Usecase.UserById(inputData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}


func (controller *UserController) Update(c echo.Context) (err error) {
	uur := new(domain.UserUpdateRequest)
	if err = c.Bind(uur); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(uur); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id, _ := strconv.ParseUint(c.Get("id").(string), 10, 32)
	inputData := input.UpdateUser{
		ID: uint(id),
		Name: uur.UserName,
	}
	user, err := controller.Usecase.Update(inputData)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
