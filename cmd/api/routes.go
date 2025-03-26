package main

import (
	"net/http"
)

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/healthcheck", app.healthcheckHandler)

	return app.recoverPanic(app.enableCORS(app.requestLogger(router)))
}
