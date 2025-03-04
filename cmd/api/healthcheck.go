package main

import (
	"net/http"
)

func (app *application) handleHealthcheck(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	err := app.writeJSON(w, jsonPayload{"api": data}, http.StatusOK, nil)
	if err != nil {
		// http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		app.serverErrorReponse(w, r, err)
	}
}
