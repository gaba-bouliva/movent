package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) router() http.Handler {
	r := chi.NewRouter()

	apiV1 := chi.NewRouter()
	apiV1.Get("/v1/healthcheck", app.handleHealthcheck)
	apiV1.Post("/v1/movies", app.handleCreateMovie)
	apiV1.Get("/v1/movies/{id}", app.handleGetMovie)
	apiV1.Get("/v1/movies", app.handleGetMovies)

	r.Mount("/api/", apiV1)

	return r
}
