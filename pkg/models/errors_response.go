package models

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct{
	Message string `json:"message"`
}

func SendError(w http.ResponseWriter, statusCode int, message string){
	w.WriteHeader(statusCode)
	resp := ErrorResponse{Message : message}
	json.NewEncoder(w).Encode(resp)
}