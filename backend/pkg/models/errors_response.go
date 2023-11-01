package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type CustomError struct {
	StatusCode    int
	Message       string
	OriginalError error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("StatusCode: %d, Message: %s, OriginalError: %v", e.StatusCode, e.Message, e.OriginalError)
}

func SendError(w http.ResponseWriter, statusCode int, message string, err error) {
	log.Printf("Error: %v", err)
	w.WriteHeader(statusCode)
	resp := ErrorResponse{Message: message}
	json.NewEncoder(w).Encode(resp)
}
