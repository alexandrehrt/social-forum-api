package routes

import (
	"github.com/go-chi/chi/v5"
	userControllers "social-api/internal/models/user/controllers"
)

func UserRoutes(r chi.Router) {
	//userGroup := app.Group("/users")
	r.Route("/users", func(r chi.Router) {
		r.Get("/", userControllers.GetUser)
		r.Post("/create", userControllers.CreateUser)
		r.Put("/{id}", userControllers.UpdateUser)
		r.Delete("/{id}", userControllers.DeleteUser)
	})
}
