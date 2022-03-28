package controllers

import (
	"net/http"
	"github.com/labstack/echo"
	"Golang-CleanArchitecture/domain"
	"Golang-CleanArchitecture/interfaces/database"
	"Golang-CleanArchitecture/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

// 依存性を定義する
func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c echo.Context) (err error) {
	ucr := new(domain.UserCreateRequest)
	if err = c.Bind(ucr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(ucr); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// 同じEmailの人がいないか確認する && UIDも
	var users domain.Users
	users, err = controller.Interactor.ExistUserByEmail(ucr.Email)
	if len(users) != 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "入力されたメールアドレスは既に登録されています。")
	}
	// DTOをUserのEntityに変換
	u := &domain.User{
		Name: ucr.UserName, 
		Email: ucr.Email,
	}
	// 同じメールアドレス、uidでerr返ってくる → 同じものを挿入したときidは進む
	user, err := controller.Interactor.Add(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) Update(c echo.Context) (err error) {
	uid := c.Get("uid").(string)
	// Jsonから構造体に変換
	uur := new(domain.UserUpdateRequest)
	if err = c.Bind(uur); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// バリデーション
	if err = c.Validate(uur); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// DTOをUserのEntityに変換
	u := &domain.User{Name: uur.UserName}
	user, err := controller.Interactor.Update(uid, u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// レスポンス
	// resopnse := domain.UserUpdateResponse{UID: user.UID, Name: user.Name, Email: user.Email}
	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) Show(c echo.Context) (err error) {
	uid := c.Get("uid").(string)
	user, err := controller.Interactor.UserByUid(uid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}