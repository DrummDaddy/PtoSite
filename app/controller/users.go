package controller

import (
	"PTOSite/app/model"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetUsers(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Получаем список всех пользователей
	users, err := model.GetAllUsers()
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
	//возвращаем список клиенту в формате JSON
	err = json.NewEncoder(rw).Encode(users)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		http.Error(rw, err.Error(), 400)
		return
	}
}
