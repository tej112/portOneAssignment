package helpers

import (
	"encoding/json"
	"net/http"
)

// WriteJSON writes the response in JSON format
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	// Set the Status code
	w.WriteHeader(status)

	// Write the JSON response
	json.NewEncoder(w).Encode(data)
}
