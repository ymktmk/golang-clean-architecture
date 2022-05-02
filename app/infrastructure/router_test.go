package infrastructure_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v9"

	"github.com/ymktmk/golang-clean-architecture/app/domain"
	"github.com/ymktmk/golang-clean-architecture/app/infrastructure"
	"github.com/ymktmk/golang-clean-architecture/app/interfaces/controllers"
)

func TestAPIRegister(t *testing.T) {
	mockDB, mock, err := infrastructure.NewDbMock()
	if err != nil {
		t.Fatal(err)
	}
	
	// mock設定
	rows := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at", "deleted_at"})
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE email = $1`)).
		WithArgs("example@gmail.com").
		WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("created_at","updated_at","deleted_at","name","email","password") VALUES ($1,$2,$3,$4,$5,$6)`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	// server
	userController := controllers.NewUserController(infrastructure.SqlMockHandler(mockDB))
	e := echo.New()
	e.Validator = &infrastructure.CustomValidator{Validator: validator.New()}
	e.POST("/register", userController.Register)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"name": "tomoki", "email": "example@gmail.com", "password": "Tomoki0901"}`)
	request, _ := http.NewRequest("POST", "/register", body)
	request.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(writer, request)

	assert.Equal(t, http.StatusOK, writer.Code)

	// response bodyの検証
	var user domain.User
	if err = json.Unmarshal(writer.Body.Bytes(), &user); err != nil {
		t.Error(err)
	}
	if user.ID != 1 && user.Name != "tomoki" && user.Email != "example@gmail.com" {
		t.Error("Cannot retrieve JSON user")
	}

	cookie, err := request.Cookie("jwt")
	if err != nil {
		return
	}
	fmt.Println(cookie)
	fmt.Println(user)
}

func TestAPITodoCreate(t *testing.T) {
	
}