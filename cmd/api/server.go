package main

import "net/http"

func (app *application) serve() error {
	server := http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	return server.ListenAndServe()
}
