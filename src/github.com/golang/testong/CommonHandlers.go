package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func getRequestBody(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	//limit request body to max 1048576 KB or 1 MB
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, v); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		return false
	}
	return true
}
