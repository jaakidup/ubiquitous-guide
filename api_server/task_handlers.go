package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// TaskUpdate ...
func TaskUpdate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// TODO update task with id
}

// TaskDelete ...
func TaskDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	//TODO delete task with id
}
