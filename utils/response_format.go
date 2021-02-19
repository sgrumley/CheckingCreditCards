package utils 

import (
	"net/http"
	"encoding/json"
)

// Define the structure of response messages
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Encode message into a http response
func Respond(w http.ResponseWriter, code int, data map[string]interface{}) {
	w.Header().Add("content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}