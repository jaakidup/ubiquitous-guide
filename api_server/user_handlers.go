package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// UserDelete ...
func (server *Server) UserDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// TODO: delete user
	// TODO Delete all the task connected to the user
	fmt.Println("Delete user with id", params.ByName("userid"))

	// delete
	delete, _ := server.db.Prepare("DELETE FROM users WHERE id=?")

	_, err := delete.Exec(params.ByName("userid"))
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}

// UserUpdate ...
func (server *Server) UserUpdate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if user.ID == 0 {
		fmt.Println("Create User")
		create, _ := server.db.Prepare("INSERT INTO users (name) VALUES (?)")
		result, err := create.Exec(user.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		id, _ := result.LastInsertId()
		w.WriteHeader(http.StatusCreated)
		response := User{
			Name: user.Name,
			ID:   int64(id),
		}
		json.NewEncoder(w).Encode(response)

	} else {
		fmt.Println("Update User")
		update, _ := server.db.Prepare("UPDATE users SET name=? where id=?")
		_, err := update.Exec(user.Name, user.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

// UsersList ...
func (server *Server) UsersList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	user := User{}
	users := []User{}
	rows, _ := server.db.Query("SELECT id, name FROM users")
	for rows.Next() {
		rows.Scan(&user.ID, &user.Name)
		users = append(users, user)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)

}
