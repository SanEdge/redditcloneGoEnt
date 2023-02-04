package routes

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/redditgoent/ent"
	"github.com/renaldyhidayatt/redditgoent/handler"
	"github.com/renaldyhidayatt/redditgoent/repository"
	"github.com/renaldyhidayatt/redditgoent/services"
)

func NewAuthRoutes(prefix string, db *ent.Client, router *chi.Mux, context context.Context) {
	repository := repository.NewAuthRepository(db, context)

	service := services.NewAuthService(repository)
	handler := handler.NewAuthHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello"))
		})

		r.Post("/login", handler.Login)

		r.Post("/register", handler.Register)
	})
}
