package database_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ymktmk/golang-clean-architecture/domain"
	"github.com/ymktmk/golang-clean-architecture/infrastructure"
	"github.com/ymktmk/golang-clean-architecture/interfaces/database"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDbMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
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
		ID: 1,
		Name: "tomoki",
		Email: "victas.tt@gmail.com",
	}
	
	// mock設定
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO users (name, email) VALUES (?, ?)`)).
		WithArgs(u.Name, u.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(1, u.Name, u.Email))
	mock.ExpectCommit()

	// SQL実行
	if err = mockDB.Create(u).Error; err != nil {
		t.Error(err)
	}
	// fmt.Println(u)
	
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("Test Create User: %v", err)
	// }

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

	u := &domain.User{
		ID: 1,
		Name: "tomoki",
		Email: "victas.tt@gmail.com",
	}

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE id = ?`)).
	WithArgs(u.ID).
	WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email"}).AddRow(u.ID, u.Name, u.Email))

	// repo := &database.UserRepository{SqlHandler: DummyHandler(mockDB)}
	// user, err := repo.FindById(int(u.ID))
	// if err != nil {
	// 	t.Fatalf("failed to find user: %s", err)
	// }
	// fmt.Println(user)

	if err = mockDB.Where("id=?", 1).First(&u).Error; err != nil {
		t.Error(err)
	}
	fmt.Println(u)
} 