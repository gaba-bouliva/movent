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
		app.notFoundResponse(w, r)
		return
	}

	movie, err := app.db.GetMovie(r.Context(), id)
	if err != nil {
		app.logger.Error(err.Error())
		// http.Error(w, "Failed to retrieve movie", http.StatusInternalServerError)
		app.serverErrorReponse(w, r, err)
		return
	}

	err = app.writeJSON(w, jsonPayload{"movie": movie}, http.StatusOK, nil)
	if err != nil {
		// http.Error(w, "Server encountered a problem and could not process your request", http.StatusInternalServerError)
		app.serverErrorReponse(w, r, err)
	}
}

func (app *application) handleGetMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.db.GetMovies(r.Context())
	if err != nil {
		app.logger.Error(err.Error())
		// http.Error(w, "error retrieving movies", http.StatusInternalServerError)
		app.serverErrorReponse(w, r, err)
		return
	}
	fmt.Printf("%+v\n", movies)
	fmt.Println("err: ", err)

	err = app.writeJSON(w, jsonPayload{"movies": movies}, http.StatusOK, nil)
	if err != nil {
		app.logger.Error(err.Error())
		// http.Error(w, "server encountered an error", http.StatusInternalServerError)
		app.serverErrorReponse(w, r, err)
	}
}
