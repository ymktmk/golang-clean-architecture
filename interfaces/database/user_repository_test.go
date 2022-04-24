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

	"gorm.io/driver/mysql"
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	// MySQL
	mockDB, err := gorm.Open(
		mysql.Dialector{
			Config: &mysql.Config{
				DriverName: "mysql",
				Conn: sqlDB,
				SkipInitializeWithVersion: true,
			},
		}, &gorm.Config{})
	return mockDB, mock, err
}

func DummyHandler(conn *gorm.DB) *infrastructure.SqlHandler {
	sqlHandler := new(infrastructure.SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func TestStore(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
	}
	
	// mock設定
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO users (created_at, updated_at, deleted_at, name, email) VALUES(?, ?, ?, ?, ?)`)).
		WithArgs(time.Now(), time.Now(), nil, "tomoki", "example@gmail.com").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	// SQL実行
	u := &domain.User{
		Name: "tomoki",
		Email: "example@gmail.com",
	}
	mockDB.Create(u)
	fmt.Println(u)

	// Repository Test
	// repo := &database.UserRepository{SqlHandler: DummyHandler(mockDB)}
	// user, err := repo.Store(u)
	// fmt.Println(user)
	// if err != nil {
	// 	t.Fatal(err)
	// }
}











func TestFindById(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
		return
	}

	var id int = 1
	// name  := "tomoki"
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE id = ?`)).
	WithArgs(id).
	WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(id))

	repo := &database.UserRepository{SqlHandler: DummyHandler(mockDB)}
	user, err := repo.FindById(id)
	if err != nil {
		t.Fatalf("failed to find user: %s", err)
	}
	fmt.Println(user)
} 