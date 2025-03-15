package main

import (
	"net/http"

	"github.com/gaba-bouliva/movent/internal/data"
	"github.com/gaba-bouliva/movent/internal/validator"
)

func (app *application) handleCreateMovie(w http.ResponseWriter, r *http.Request) {
	var reqBody struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}
	// get the body
	err := app.readJSON(w, r, &reqBody)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// validate the data in body

	movie := data.Movie{
		Title:   reqBody.Title,
		Year:    reqBody.Year,
		Runtime: reqBody.Runtime,
		Genres:  reqBody.Genres,
	}
	v := validator.New()
	data.ValidateMovie(v, &movie)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	createMoveParams := data.CreateMovieParams{
		Title:   reqBody.Title,
		Year:    reqBody.Year,
		Runtime: reqBody.Runtime,
		Genres:  reqBody.Genres,
	}

	_, err = app.db.CreateMovie(r.Context(), createMoveParams)
	if err != nil {
		app.serverErrorReponse(w, r, err)
		return
	}
	err = app.writeJSON(w, jsonPayload{"message": "movie created successfully"}, http.StatusCreated, nil)
	if err != nil {
		app.logError(r, err)
	}

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

	err = app.writeJSON(w, jsonPayload{"movies": movies}, http.StatusOK, nil)
	if err != nil {
		app.logger.Error(err.Error())
		// http.Error(w, "server encountered an error", http.StatusInternalServerError)
		app.serverErrorReponse(w, r, err)
	}
}

func (app *application) handleUpdateMovie(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie, err := app.db.GetMovie(r.Context(), id)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var reqBody struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}
	err = app.readJSON(w, r, &reqBody)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	movie.Title = reqBody.Title
	movie.Year = reqBody.Year
	movie.Runtime = reqBody.Runtime
	movie.Genres = reqBody.Genres

	v := validator.New()
	data.ValidateMovie(v, &movie)

	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	updateMovieParams := data.UpdateMovieParams{
		Title:   movie.Title,
		Year:    movie.Year,
		Runtime: movie.Runtime,
		Genres:  movie.Genres,
	}

	err = app.db.UpdateMovie(r.Context(), updateMovieParams)
	if err != nil {
		app.serverErrorReponse(w, r, err)
		return
	}

	err = app.writeJSON(w, jsonPayload{"message": "movie updated successfully"}, http.StatusAccepted, nil)
	if err != nil {
		app.serverErrorReponse(w, r, err)
	}
}

func (app *application) handleDeleteMovie(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	_, err = app.db.GetMovie(r.Context(), id)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.db.DeleteMovie(r.Context(), id)
	if err != nil {
		app.serverErrorReponse(w, r, err)
		return
	}

	err = app.writeJSON(w, jsonPayload{"message": "movie deleted successfully"}, http.StatusNoContent, nil)
	if err != nil {
		app.serverErrorReponse(w, r, err)
		return
	}
}
