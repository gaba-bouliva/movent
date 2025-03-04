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
		app.serverErrorReponse(w, r, err)
	}
}
