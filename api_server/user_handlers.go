package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// UserDelete ...
func UserDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// TODO: delete user

}

// UserTasks ...
func UserTasks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(params.ByName("userid")))
	// TODO get all the user tasks

}

// UserUpdate ...
func UserUpdate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// TODO update user

}

// UsersList ...
func UsersList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(params.ByName("userid")))

	// TODO fetch all users

}
