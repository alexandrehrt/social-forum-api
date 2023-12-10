package routes

import (
	"github.com/go-chi/chi/v5"
	groupControllers "social-api/internal/models/group/controllers"
)

func GroupRoutes(r chi.Router) {
	r.Route("/groups", func(r chi.Router) {
		//r.Get("/", getAllGroups)
		//r.Get("/{id}", getGroup)
		r.Post("/create", groupControllers.CreateGroup)
		//r.Put("/{id}", updateGroup)
		//r.Delete("/{id}", deleteGroup)
	})
}
