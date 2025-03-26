package main

import (
	"fmt"
	"log"
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

	log.Printf("starting API %s server on %s", app.config.env, server.Addr)
	return server.ListenAndServe()
}
