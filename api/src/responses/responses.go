package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON return a default response to request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if error := json.NewEncoder(w).Encode(data); error != nil {
		log.Fatal(error)
	}
}

// Error return an error in JSON format
func Error(w http.ResponseWriter, statusCode int, error error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: error.Error(),
	})
}
