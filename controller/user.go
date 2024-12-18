package controller

import (
	"encoding/json"
	"net/http"
)

func Register(respw http.ResponseWriter, req *http.Request) {
	
}

func Login(respw http.ResponseWriter, req *http.Request) {
	
}

func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"message": message,
	})
}