package models

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct{
	Message string `json:"message"`
}

func SendError(w http.ResponseWriter, statusCode int, message string, err error){
	log.Printf("Error: %v", err)
	w.WriteHeader(statusCode)
	resp := ErrorResponse{Message : message}
	json.NewEncoder(w).Encode(resp)
}