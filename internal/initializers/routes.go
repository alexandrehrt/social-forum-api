package initializers

import (
	"github.com/go-chi/chi/v5"
	"social-api/internal/routes"
)

func Routes(r chi.Router) {
	routes.UserRoutes(r)
	//docs.GroupRoutes(r)
	//docs.Login(r)
}
