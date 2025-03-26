package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) serve() error {
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.logger.LogInfo("starting API server", map[string]string{
		"address":     server.Addr,
		"environment": app.config.env,
	})
	return server.ListenAndServe()
}
