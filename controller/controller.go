package controller

import (
	"encoding/json"
	"net/http"
)

// GetSaludo .
func GetSaludo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Hola desde API ")
}
