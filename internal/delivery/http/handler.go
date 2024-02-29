package delivery

import (
	"log/slog"
	"tasktracker/internal/service"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services *service.Services

	log slog.Logger
}

// TODO : Add middleware
func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/users", func(r chi.Router) {
		r.Post("/sign-up", h.SignUp)
		r.Post("/sign-in", h.SignIn)
	})

	// TODO : Init task tracker routes

	// TODO : Return Handler

	return router
}

func NewHandler(services *service.Services, log slog.Logger) *Handler {
	return &Handler{
		services: services,
		log:      log,
	}
}
