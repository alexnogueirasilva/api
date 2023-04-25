package response

import (
	"encoding/json"
	"net/http"
)

// JSON returns a JSON response to the request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			Error(w, http.StatusInternalServerError, err)
			return
		}
	}
}

// Error returns an error message to the request
func Error(w http.ResponseWriter, statusCode int, err error) {

	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
