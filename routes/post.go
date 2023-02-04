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

func NewPostRoutes(prefix string, db *ent.Client, router *chi.Mux, context context.Context) {
	repository := repository.NewPostRepository(db, context)

	service := services.NewPostService(repository)
	handler := handler.NewPostHandler(service)

	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello"))
		})
		r.Post("/", handler.Create)
		r.Get("/:id", handler.FindById)
		r.Get("/", handler.FindAll)
		r.Get("/subreddit/:id", handler.FindPostBySubreddit)
		r.Get("/findbyemail", handler.FindByEmail)
	})
}
