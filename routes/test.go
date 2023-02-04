package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/renaldyhidayatt/redditgoent/dto/response"
	"github.com/renaldyhidayatt/redditgoent/middlewares"
)

func NewTestRoutes(prefix string, router *chi.Mux) {
	router.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuthentication)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			token := r.Context().Value("decoded")
			tokenStr := token.(string)
			var decodedToken map[string]interface{}
			json.Unmarshal([]byte(tokenStr), &decodedToken)
			email, ok := decodedToken["sub"].(string)
			if !ok {
				response.ResponseError(w, http.StatusBadRequest, fmt.Errorf("failed convert"))
			}

			w.Write([]byte(email))

		})

	})
}
