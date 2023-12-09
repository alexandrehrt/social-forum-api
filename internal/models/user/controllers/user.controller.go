package userControllers

import (
	"encoding/json"
	"net/http"
	"social-api/internal/entity"
	userUseCases "social-api/internal/models/user/usecases"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get User"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := new(entity.User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	userResponse, err := userUseCases.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(userResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update User " + id))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Delete User " + id))
}
