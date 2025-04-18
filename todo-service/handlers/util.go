package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// respondInternalServerError writes an internal server error JSON response.
func respondInternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	resp := map[string]string{
		"message": err.Error(),
	}
	jsonResp, jsonErr := json.Marshal(resp)
	if jsonErr != nil {
		// Log the error since writing JSON failed.
		log.Fatalf("Error during JSON marshal: %s", jsonErr)
	}

	_, writeErr := w.Write(jsonResp)
	if writeErr != nil {
		log.Printf("Error writing response: %s", writeErr)
	}
}
