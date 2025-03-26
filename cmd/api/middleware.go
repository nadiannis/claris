package main

import (
	"net/http"
)

func (app *application) requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.logger.LogInfo("request received", map[string]string{
			"remote_addr": r.RemoteAddr,
			"proto":       r.Proto,
			"method":      r.Method,
			"url":         r.URL.RequestURI(),
		})

		next.ServeHTTP(w, r)
	})
}
