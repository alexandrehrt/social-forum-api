package userControllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"social-api/internal/entity"
	userUseCases "social-api/internal/models/user/usecases"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := userUseCases.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var userResponses []*userUseCases.UserResponse
	for _, user := range users {
		userResponse := userUseCases.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}
		userResponses = append(userResponses, &userResponse)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(userResponses)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	userResponse, err := userUseCases.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(userResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user *entity.User
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
