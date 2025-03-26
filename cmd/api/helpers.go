package main

import (
	"encoding/json"
	"net/http"

	"github.com/nadiannis/claris/internal/types"
)

func (app *application) writeJSON(w http.ResponseWriter, status int, data types.Envelope, headers http.Header) error {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	jsonBytes = append(jsonBytes, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonBytes)

	return nil
}
