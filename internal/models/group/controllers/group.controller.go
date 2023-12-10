package groupControllers

import (
	"encoding/json"
	"net/http"
	"social-api/internal/entity"
	groupUsecases "social-api/internal/models/group/usecases"
)

//func GetAllGroups(w http.ResponseWriter, r *http.Request) {
//	users, err := userUseCases.GetAllUsers()
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		w.Write([]byte(err.Error()))
//		return
//	}
//
//	var userResponses []*userDTO.UserResponse
//	for _, user := range users {
//		userResponse := userDTO.UserResponse{
//			ID:       user.ID,
//			Username: user.Username,
//			Email:    user.Email,
//		}
//		userResponses = append(userResponses, &userResponse)
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	err = json.NewEncoder(w).Encode(userResponses)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		w.Write([]byte(err.Error()))
//	}
//}
//
//func GetGroup(w http.ResponseWriter, r *http.Request) {
//	id := chi.URLParam(r, "id")
//
//	userResponse, err := userUseCases.GetUser(id)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		w.Write([]byte(err.Error()))
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//	err = json.NewEncoder(w).Encode(userResponse)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		w.Write([]byte(err.Error()))
//	}
//
//}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	group := new(entity.Group)
	err := json.NewDecoder(r.Body).Decode(group)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	createGroupResponse, err := groupUsecases.Create(group)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(createGroupResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//func UpdateGroup(w http.ResponseWriter, r *http.Request) {
//	id := r.URL.Query().Get("id")
//
//	w.WriteHeader(http.StatusOK)
//	w.Write([]byte("Update User " + id))
//}
//
//func DeleteGroup(w http.ResponseWriter, r *http.Request) {
//	id := chi.URLParam(r, "id")
//
//	err := userUseCases.DeleteUser(id)
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		w.Write([]byte(err.Error()))
//		return
//	}
//
//	w.WriteHeader(http.StatusNoContent)
//}
