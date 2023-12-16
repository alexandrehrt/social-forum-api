package userControllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"social-api/internal/entity"
	"social-api/internal/models/user/dtos"
	userUseCases "social-api/internal/models/user/usecases"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, appError := userUseCases.GetAllUsers()
	if appError != nil {
		w.WriteHeader(appError.StatusCode)
		w.Write([]byte(appError.Message))
		return
	}

	var userResponses []*userDTO.UserResponse
	for _, user := range users {
		userResponse := userDTO.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}
		userResponses = append(userResponses, &userResponse)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(userResponses)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}

func FindUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	userResponse, appErr := userUseCases.FindUserByID(id)
	if appErr != nil {
		w.WriteHeader(appErr.StatusCode)
		w.Write([]byte(appErr.Message))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(userResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := new(entity.User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	appError := userUseCases.Create(user)
	if appError != nil {
		w.WriteHeader(appError.StatusCode)
		w.Write([]byte(appError.Message))
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Update User " + id))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	appError := userUseCases.DeleteUser(id)
	if appError != nil {
		w.WriteHeader(appError.StatusCode)
		w.Write([]byte(appError.Message))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
