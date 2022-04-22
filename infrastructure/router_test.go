package infrastructure

import (
	// "encoding/json"
	// "net/http"
	// "net/http/httptest"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	// "github.com/ymktmk/golang-clean-architecture/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

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

func TestUserGet(t *testing.T) {

	// mockDB, mock, err := NewDbMock()
	// if err != nil {       
	// 	t.Fatal(err)
	// }
	
	// e := NewRouter()
	// writer := httptest.NewRecorder()
	// request, _ := http.NewRequest("GET", "/user", nil)
	
	// e.ServeHTTP(writer, request)
	
	// if writer.Code != 200 {
	// 	t.Errorf("Respose Code is %v", writer.Code)
	// }
	// var user domain.User
	// json.Unmarshal(writer.Body.Bytes(), &user)
	// if user.ID != 1 {
	// 	t.Error("Cannot retrieve JSON user")
	// }
}