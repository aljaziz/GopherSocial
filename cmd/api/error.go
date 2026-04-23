package main

import (
	"log"
	"net/http"
)

func (app *application) internalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	errorJSON(w, http.StatusInternalServerError, "The server encountered a problem")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad Request: %s path: %s error: %s", r.Method, r.URL.Path, err)

	errorJSON(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Not found error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	errorJSON(w, http.StatusNotFound, "Not found")
}

func (app *application) forbiddenResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Forbidden: %s path: %s error: %s", r.Method, r.URL.Path, err)

	errorJSON(w, http.StatusForbidden, "Forbidden")
}

func (app *application) unauthorizedResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Unauthorized error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	errorJSON(w, http.StatusUnauthorized, "Unauthorized")
}
