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
	"github.com/ymktmk/golang-clean-architecture/domain"
	"github.com/ymktmk/golang-clean-architecture/infrastructure"
	"github.com/ymktmk/golang-clean-architecture/interfaces/controllers"
	"github.com/ymktmk/golang-clean-architecture/interfaces/database"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	return mockDB, mock, err
}

func DummyHandler(conn *gorm.DB) database.SqlHandler {
	sqlHandler := new(infrastructure.SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func TestCreate(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {       
		t.Fatal(err)
	}
	
	// server 
	userController := controllers.NewUserController(DummyHandler(mockDB))
	e := echo.New()
	e.Validator = &infrastructure.CustomValidator{Validator: validator.New()}
	e.POST("/users/create", userController.Create)
	
	// mock設定
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("created_at","updated_at","deleted_at","name","email") VALUES ($1,$2,$3,$4,$5)`)).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	// recorder
	writer := httptest.NewRecorder()

	// request
	payload := strings.NewReader(`{"name": "tomoki", "email": "tt@gmail.com"}`)
	request, _ := http.NewRequest("POST", "/users/create", payload)
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
	mockDB, mock, err := NewDbMock()
	if err != nil {       
		t.Fatal(err)
	}

	userController := controllers.NewUserController(DummyHandler(mockDB))
	e := echo.New()
	e.GET("/user", userController.Show)

	// mock設定
	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at", "deleted_at"}).
	AddRow(1, "tomoki", "example@gmail.com", time.Now(), time.Now(), nil)
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE id = $1`)).
		WithArgs(1).
		WillReturnRows(rows)

	// recorder
	writer := httptest.NewRecorder()
	
	// request
	request, _ := http.NewRequest("GET", "/user", nil)
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