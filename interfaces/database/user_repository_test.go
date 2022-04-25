package database_test

import (
	"fmt"
	"regexp"
	"testing"
	"time"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ymktmk/golang-clean-architecture/domain"
	"github.com/ymktmk/golang-clean-architecture/infrastructure"
	"github.com/ymktmk/golang-clean-architecture/interfaces/database"
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

func TestStore(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	
	u := &domain.User{
		Name: "sheep",
		Email: "example@gmail.com",
	}
	
	// mock設定
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("created_at","updated_at","deleted_at","name","email") VALUES ($1,$2,$3,$4,$5)`)).
		WillReturnRows(rows)
	mock.ExpectCommit()

	repo := &database.UserRepository{SqlHandler: DummyHandler(mockDB)}
	user, err := repo.Store(u)
	fmt.Println(user)
	if err != nil {
		t.Fatal(err)
	}
	
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Create User: %v", err)
	}
}


func TestFindById(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	// mock設定
	rows := sqlmock.NewRows([]string{"id", "name", "email", "created_at", "updated_at", "deleted_at"}).
	AddRow(1, "tomoki", "example@gmail.com", time.Now(), time.Now(), nil)
	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "users" WHERE id = $1`)).
		WithArgs(1).
		WillReturnRows(rows)

	repo := &database.UserRepository{SqlHandler: DummyHandler(mockDB)}
	user, err := repo.FindById(1)
	fmt.Println(user)
	if err != nil {
		t.Fatal(err)
	}
}