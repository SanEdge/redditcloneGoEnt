package routes

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/redditgoent/ent"
	"github.com/renaldyhidayatt/redditgoent/handler"
	"github.com/renaldyhidayatt/redditgoent/middlewares"
	"github.com/renaldyhidayatt/redditgoent/repository"
	"github.com/renaldyhidayatt/redditgoent/services"
)

func NewCommentRoutes(prefix string, db *ent.Client, router *chi.Mux, context context.Context) {
	repository := repository.NewCommentRepository(db, context)

	service := services.NewCommentService(repository)
	handler := handler.NewCommentHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello"))
		})
		r.Post("/", handler.Create)
		r.Get("/post", handler.FindAllCommentForPost)
		r.Get("/user", handler.FindAllCommentForUser)

	})
}
