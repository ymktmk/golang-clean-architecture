package infrastructure_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/ymktmk/golang-clean-architecture/domain"
	"github.com/ymktmk/golang-clean-architecture/domain/gorm"
	"github.com/ymktmk/golang-clean-architecture/infrastructure"
	"github.com/ymktmk/golang-clean-architecture/interfaces/controllers"
	"github.com/ymktmk/golang-clean-architecture/utils"
	"gopkg.in/go-playground/validator.v9"
)

func TestCreate(t *testing.T) {
	mockDB, mock, err := utils.NewDbMock()
	if err != nil {       
		t.Fatal(err)
	}
	
	// mock設定
	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at", "deleted_at"})
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE email = $1`)).
		WithArgs("tt@gmail.com").
		WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("created_at","updated_at","deleted_at","name","email") VALUES ($1,$2,$3,$4,$5)`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()
	
	// server 
	userController := controllers.NewUserController(utils.DummyHandler(mockDB))
	e := echo.New()
	e.Validator = &infrastructure.CustomValidator{Validator: validator.New()}
	e.POST("/users/create", userController.Create)

	writer := httptest.NewRecorder()
	body := strings.NewReader(`{"name": "tomoki", "email": "tt@gmail.com"}`)
	request, _ := http.NewRequest("POST", "/users/create", body)
	request.Header.Set("Content-Type", "application/json")
	e.ServeHTTP(writer, request)
	
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	// response bodyの検証
	var user domain.User
	json.Unmarshal(writer.Body.Bytes(), &user)
	if user.ID != 1 {
		t.Error("Cannot retrieve JSON user")
	}

	fmt.Println(user)
}

func TestShow(t *testing.T) {
	mockDB, mock, err := utils.NewDbMock()
	if err != nil {       
		t.Fatal(err)
	}

	// response structure definition
	user := &domain.User{
		Model: gorm.Model{
			ID: 1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: time.Time{},
		},
		Name: "tomoki",
		Email: "example@gmail.com",
	}

	user_json, _ := json.Marshal(user)
	
	// mock設定
	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at", "deleted_at"}).
	AddRow(user.ID, user.Name, user.Email, user.CreatedAt, user.UpdatedAt, user.DeletedAt)
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE id = $1`)).
		WithArgs(user.ID).
		WillReturnRows(rows)

      // server
	userController := controllers.NewUserController(utils.DummyHandler(mockDB))
	e := echo.New()
	e.GET("/user", userController.Show)

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/user", nil)
	e.ServeHTTP(writer, request)
	
	// status code validation
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}

	// json response validation
	assert.JSONEq(t, string(user_json), string(writer.Body.Bytes()))
}