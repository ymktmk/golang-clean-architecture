package controllers

import (
	"Golang-CleanArchitecture/domain"
	"encoding/json"
	"fmt"
	"net/http"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {

	// Request
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var user domain.User
	json.Unmarshal(body, &user)

	// この間のビジネスロジック
	// DBに保存するなど

	// Response
	response, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

