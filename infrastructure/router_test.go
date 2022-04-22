package infrastructure

import (
	// "net/http"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ymktmk/golang-clean-architecture/domain"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func TestUserGet(t *testing.T) {
	e := NewRouter()
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/user", nil)
	
	e.ServeHTTP(writer, request)
	
	// 結果
	if writer.Code != 200 {
		t.Errorf("Respose Code is %v", writer.Code)
	}
	var user domain.User
	json.Unmarshal(writer.Body.Bytes(), &user)
	if user.ID != 1 {
		t.Error("Cannot retrieve JSON user")
	}
}