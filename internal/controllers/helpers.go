package controllers

import (
	"encoding/json"
	"net/http"
)

const (
	XUserId = "X-UserId"
)

func Error(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func Success(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(message)
}
