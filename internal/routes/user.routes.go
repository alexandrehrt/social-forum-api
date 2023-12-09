package routes

import (
	"github.com/go-chi/chi/v5"
	userControllers "social-api/internal/models/user/controllers"
)

func UserRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Get("/", userControllers.GetAllUsers)
		r.Get("/{id}", userControllers.GetUser)
		r.Post("/create", userControllers.CreateUser)
		r.Put("/{id}", userControllers.UpdateUser)
		r.Delete("/{id}", userControllers.DeleteUser)
	})
}
