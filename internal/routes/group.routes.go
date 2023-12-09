package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GroupRoutes(r chi.Router) {
	r.Route("/groups", func(r chi.Router) {
		r.Get("/", getGroup)
		r.Post("/", createGroup)
		r.Put("/{id}", updateGroup)
		r.Delete("/{id}", deleteGroup)
	})
}

func getGroup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get Group"))
}

func createGroup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create Group"))
}

func updateGroup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update Group"))
}

func deleteGroup(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete Group"))
}
