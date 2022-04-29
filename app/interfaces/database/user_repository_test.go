package database_test

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/ymktmk/golang-clean-architecture/app/domain"
	"github.com/ymktmk/golang-clean-architecture/app/interfaces/database"
	"github.com/ymktmk/golang-clean-architecture/app/utils"
)

func TestStore(t *testing.T) {
	mockDB, mock, err := utils.NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}

	u := &domain.User{
		Name:  "sheep",
		Email: "example@gmail.com",
	}

	// mock設定
	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("created_at","updated_at","deleted_at","name","email") VALUES ($1,$2,$3,$4,$5)`)).
		WillReturnRows(rows)
	mock.ExpectCommit()

	// repository 初期化
	repo := &database.UserRepository{SqlHandler: utils.SqlMockHandler(mockDB)}
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
	mockDB, mock, err := utils.NewDbMock()
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

		// repository 初期化
	repo := &database.UserRepository{SqlHandler: utils.SqlMockHandler(mockDB)}
	user, err := repo.FindById(1)
	fmt.Println(user)
	if err != nil {
		t.Fatal(err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Test Find User: %v", err)
	}
}
