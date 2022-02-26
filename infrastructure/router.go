package infrastructure

import (
	"github.com/gorilla/mux"
	"Golang-CleanArchitecture/interfaces/controllers"
)

func Routing() *mux.Router {
	r := mux.NewRouter()
	// SQLHandlerをInterfaces層のControllerに
	userController := controllers.NewUserController(NewSqlHandler())
	r.HandleFunc("/user/create", userController.Create).Methods("POST")
	return r
}