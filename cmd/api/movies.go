package main

import (
	"fmt"
	"net/http"
)

func (app *application) handleCreateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) handleGetMovie(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of movie %d\n", id)
}

func (app *application) handleGetMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.db.GetMovies(r.Context())
	if err != nil {
		http.Error(w, "error retrieving movies", http.StatusInternalServerError)
		return
	}

	app.writeJSON(w, jsonPayload{"movies": movies}, http.StatusOK, nil)
}
