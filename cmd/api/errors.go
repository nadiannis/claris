package main

import (
	"net/http"

	"github.com/nadiannis/claris/internal/types"
)

func (app *application) logError(r *http.Request, err error) {
	app.logger.LogError(err, map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	data := types.Envelope{"error": message}

	err := app.writeJSON(w, status, data, nil)
	if err != nil {
		app.logError(r, err)

		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)

	message := "server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}
