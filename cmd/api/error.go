package main

import (
	"net/http"
)

func (app *application) internalServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorw("Internal Server Error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	errorJSON(w, http.StatusInternalServerError, "The server encountered a problem")
}

func (app *application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("Bad Request", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	errorJSON(w, http.StatusBadRequest, err.Error())
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("Not Found Error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	errorJSON(w, http.StatusNotFound, "Not found")
}

func (app *application) forbiddenResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnw("Forbidden", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	errorJSON(w, http.StatusForbidden, "Forbidden")
}

func (app *application) unauthorizedResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("Unauthorized Error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	errorJSON(w, http.StatusUnauthorized, "Unauthorized")
}

func (app *application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Errorf("Conflict Response", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	errorJSON(w, http.StatusConflict, err.Error())
}

func (app *application) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Warnf("Unauthorized Basic Error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	w.Header().Set("WWW-Authenticate", `Basic realm="Restricted", charset="UTF-8"`)

	writeJSON(w, http.StatusUnauthorized, "unauthorized")
}
