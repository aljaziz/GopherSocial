package main

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_046_578 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}

func errorJSON(w http.ResponseWriter, status int, message string) error {
	type jsonResponse struct {
		Error string `json:"error"`
	}
	return writeJSON(w, status, &jsonResponse{Error: message})
}

func jsonResponse(w http.ResponseWriter, status int, data any) error {
	type jsonResponse struct {
		Data any `json:"data"`
	}
	return writeJSON(w, status, &jsonResponse{Data: data})
}
