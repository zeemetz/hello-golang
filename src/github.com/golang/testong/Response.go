package main

import (
	"encoding/json"
	"net/http"
)

//ErrorResponse is
type ErrorResponse struct {
	ACK     bool
	Message string
}

func responseError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusPreconditionFailed)
	errorResponse := ErrorResponse{ACK: false, Message: message}
	if err2 := json.NewEncoder(w).Encode(errorResponse); err2 != nil {
		panic(err2)
	}
}

func responseSuccess(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		panic(err)
	}
}
