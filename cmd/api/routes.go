package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/nadiannis/claris/internal/types"
)

func (app *application) routes() http.Handler {
	router := http.NewServeMux()

	router.Handle("GET /api/v1", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsonBytes, err := json.Marshal(types.Envelope{"message": "Hello, World!"})
		if err != nil {
			log.Printf("err: %v, request_method: %s, request_url: %s", err, r.Method, r.URL.String())

			jsonBytes, err := json.Marshal(types.Envelope{"error": err.Error()})
			if err != nil {
				log.Printf("err: %v, request_method: %s, request_url: %s", err, r.Method, r.URL.String())

				w.WriteHeader(http.StatusInternalServerError)

				return
			}

			jsonBytes = append(jsonBytes, '\n')

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonBytes)

			return
		}

		jsonBytes = append(jsonBytes, '\n')

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}))

	return router
}
