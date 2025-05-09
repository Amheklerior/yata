package utils

import (
	"encoding/json"
	"net/http"
)

// The data to send as JSON
type Envelope map[string]any

func WriteJSON(w http.ResponseWriter, status int, data Envelope) error {
	// parse data into (indented) json format
	js, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	// append a new line char at the end
	js = append(js, '\n')

	// write json response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
