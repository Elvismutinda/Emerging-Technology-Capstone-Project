package commonutils

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func setResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Request-Id", uuid.NewString())
	w.Header().Set("Content-Security-Policy", "default-src 'self';")
	w.Header().Set("Strict-Transport-Security", "max-age="+"10000000000000")
}

func HTTPResponse(w http.ResponseWriter, response Response, statusCode int) error {
	setResponseHeaders(w)
	w.WriteHeader(statusCode)
	// Write response
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return err
	}

	return nil
}
