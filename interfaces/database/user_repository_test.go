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
	
	// 初期化できない
	// mockDB, err := gorm.Open(mysql.New(mysql.Config{
	// 	Conn: sqlDB,
	// }), &gorm.Config{})
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
		return
	}
	
	u := &domain.User{
		CreatedAt: &time.Time{},
		UpdatedAt: &time.Time{},
		DeletedAt: nil,
		Name: "tomoki",
		Email: "victas.tt@gmail.com",
	}

	// mock設定
	// `INSERT INTO "users" ("name", "email") VALUES ($1, $2)`
	// mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("created_at", "updated_at", "deleted_at", "name", "email") VALUES ($1, $2, $3, $4, $5)`)).
		// クエリパラメータの指定
		WithArgs(time.Now, time.Now, nil, u.Name, u.Email).
		WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow(u.Name))
	// mock.ExpectCommit()
	
	repo := &database.UserRepository{SqlHandler: DummyHandler(mockDB)}
	user, err := repo.Store(u)
	// if err != nil {
	// 	t.Fatalf("failed to create user: %s", err)
	// }
	fmt.Println(user)
}

func TestFindById(t *testing.T) {
	mockDB, mock, err := NewDbMock()
	if err != nil {
		t.Errorf("Failed to initialize mock DB: %v", err)
		return
	}
	repo := &database.UserRepository{SqlHandler: DummyHandler(mockDB)}
	
	var id int = 1
	var name string = "tomoki"
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE (id = $1)`)).
	// クエリパラメータの指定
	WithArgs(id).
	WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(id, name))

	user, err := repo.FindById(id)
	// if err != nil {
	// 	t.Fatalf("failed to find user: %s", err)
	// }
	fmt.Println(user)
} 