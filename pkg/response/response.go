package response

import (
	"encoding/json"
	"net/http"
)

// ErrorMessage standarized error response.
type ErrorMessage struct {
	Message string `json:"message"`
}

// Map is a convenient way to create objects of unknown types.
type Map map[string]interface{}

// JSON standarized JSON response.
func JSON(w http.ResponseWriter, r *http.Request, statusCode int, data interface{}) error {
	if data == nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(statusCode)
		return nil
	}

	j, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	w.Write(j)
	return nil
}

// HTTPError standarized error response in JSON format.
func HTTPError(w http.ResponseWriter, r *http.Request, statusCode int, message string) error {
	msg := ErrorMessage{
		Message: message,
	}

	return JSON(w, r, statusCode, msg)
}
