package routes

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/redditgoent/ent"
	"github.com/renaldyhidayatt/redditgoent/handler"
	"github.com/renaldyhidayatt/redditgoent/middlewares"
	"github.com/renaldyhidayatt/redditgoent/repository"
	"github.com/renaldyhidayatt/redditgoent/services"
)

func NewSubRedditRoutes(prefix string, db *ent.Client, router *chi.Mux, context context.Context) {
	repository := repository.NewSubRedditRepository(db, context)

	service := services.NewSubRedditService(repository)
	handler := handler.NewSubRedditHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)
		r.Post("/", handler.Create)
		r.Get("/", handler.FindAll)
		r.Get("/:id", handler.FindById)
	})
}
