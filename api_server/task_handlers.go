package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// TaskUpdate ...
func (server *Server) TaskUpdate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	fmt.Println(task)

	if task.ID == 0 {
		fmt.Println("Create Task")
		create, _ := server.db.Prepare("INSERT INTO tasks (description, state, user_id) VALUES (?, ?, ?)")
		result, err := create.Exec(task.Description, task.State, params.ByName("userid"))
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		id, _ := result.LastInsertId()
		response := Task{
			ID:          int64(id),
			Description: task.Description,
			State:       task.State,
			UserID:      task.UserID,
		}
		// w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	} else {
		fmt.Println("Update Task")
		update, _ := server.db.Prepare("UPDATE tasks SET description=?, state=? WHERE id=?")
		_, err := update.Exec(task.Description, task.State, task.ID)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}

}

// TaskDelete ...
func (server *Server) TaskDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	delete, _ := server.db.Prepare("DELETE FROM tasks WHERE user_id=? and id=?")

	_, err := delete.Exec(params.ByName("userid"), params.ByName("taskid"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// UserTasks ...
func (server *Server) UserTasks(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	task := Task{}
	tasks := []Task{}
	query, _ := server.db.Prepare("SELECT id, description, state FROM tasks WHERE user_id=?")
	rows, err := query.Query(params.ByName("userid"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for rows.Next() {
		rows.Scan(&task.ID, &task.Description, task.State)
		tasks = append(tasks, task)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&tasks)
}
