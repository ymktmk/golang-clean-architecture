// リクエストを受け取るところ
package infrastructure

import (
	"github.com/gorilla/mux"
	"Golang-CleanArchitecture/interfaces/controllers"
)

func Routing() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user/create", controllers.UserCreate).Methods("POST")
	return r
}