package delivery

import (
	"log/slog"
	"tasktracker/internal/service"
	"tasktracker/pkg/auth"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services *service.Services

	tokenManager auth.Manager
	log          slog.Logger
}

// TODO : Add middleware
func (h *Handler) InitRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Route("/users", func(r chi.Router) {
		r.Post("/sign-up", h.SignUp)
		r.Post("/sign-in", h.SignIn)
		r.Post("/auth/refresh", h.Refresh)
	})

	router.Route("/tasks", func(r chi.Router) {
		r.Use(h.tokenManager.MiddlewareJWT)
		r.Get("/ping", h.Ping)
	})

	// TODO : Init task tracker routes

	// TODO : Return Handler

	return router
}

func NewHandler(services *service.Services, tokenManager auth.Manager, log slog.Logger) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
		log:          log,
	}
}
