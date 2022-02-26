package controllers

import (
	"Golang-CleanArchitecture/domain"
	"Golang-CleanArchitecture/usecase"
	"Golang-CleanArchitecture/interfaces/database"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserController struct {
    Interactor usecase.UserInteractor
}

// これがわからん
func NewUserController(sqlHandler database.SqlHandler) *UserController {
    return &UserController{
        Interactor: usecase.UserInteractor{
            UserRepository: &database.UserRepository{
                SqlHandler: sqlHandler,
            },
        },
    }
}

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {

	// convert to json
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var user domain.User
	json.Unmarshal(body, &user)

	// usecase
	user, err := controller.Interactor.Add(user)
    if err != nil {
        fmt.Println(err)
	}
	
	// response
	fmt.Println(user)

	// response, err := json.Marshal(user)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.Write(response)
}

