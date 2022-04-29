package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"

	"github.com/ymktmk/golang-clean-architecture/app/domain"
	"github.com/ymktmk/golang-clean-architecture/app/interfaces/database"
	"github.com/ymktmk/golang-clean-architecture/app/usecase"
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

func (controller *UserController) Register(c echo.Context) (err error) {
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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// パスワードをハッシュ化
	password, _ := bcrypt.GenerateFromPassword([]byte(ucr.Password), 10)
	// DTOをUserのEntityに変換
	u := &domain.User{
		Name:     ucr.UserName,
		Email:    ucr.Email,
		Password: password,
	}
	// 同じメールアドレス、uidでerr返ってくる → 同じものを挿入したときidは進む
	user, err := controller.Interactor.Add(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// // jwt tokenの検証
	// token := utils.CreateToken(int(user.ID))
	// // Cookie
	// cookie := http.Cookie{
	// 	Name:     "jwt",
	// 	Value:    token,
	// 	Expires:  time.Now().Add(time.Hour * 24),
	// 	HttpOnly: true,
	// }
	// c.SetCookie(&cookie)

	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) Show(c echo.Context) (err error) {
	user, err := controller.Interactor.UserById(1)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
