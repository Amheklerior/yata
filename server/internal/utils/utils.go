package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Amheklerior/yata/server/internal/store"
	"github.com/go-chi/chi/v5"
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

func GetTaskIdFromURLParam(r *http.Request) (store.TaskId, error) {
	taskIdUrlParam := chi.URLParam(r, "id")
	if taskIdUrlParam == "" {
		return 0, fmt.Errorf("invalid Id: '%v'", taskIdUrlParam)
	}

	taskId, err := strconv.ParseInt(taskIdUrlParam, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid Id. '%v' is not an integer value", taskIdUrlParam)
	}

	return store.TaskId(taskId), nil
}
