package initializers

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"social-api/internal/routes"
)

func Routes(r chi.Router) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	routes.UserRoutes(r)
	routes.GroupRoutes(r)
	//routes.Login(r)
}
