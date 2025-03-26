package main

import (
	"net/http"

	"github.com/nadiannis/claris/internal/types"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	data := types.Envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     types.APIVersion,
		},
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
