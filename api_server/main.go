package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	_ "github.com/mattn/go-sqlite3"
)

// Server ...
type Server struct {
	router *httprouter.Router
	db     *sql.DB
}

var server *Server

func setup() {
	server = &Server{
		router: httprouter.New(),
	}

	var err error
	server.db, err = sql.Open("sqlite3", "./todos.sqlite")
	if err != nil {
		fmt.Println(err.Error())
	}
	userTable, _ := server.db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT)")
	userTable.Exec()
	taskTable, _ := server.db.Prepare("CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, description TEXT, state TEXT, user_id TEXT)")
	taskTable.Exec()

	// TODO insert dummy data
}

func routes() {
	server.router.GET("/user", Logger(server.UsersList))
	server.router.POST("/user", Logger(server.UserUpdate))
	server.router.DELETE("/user/:userid", Logger(server.UserDelete))

	server.router.GET("/user/:userid", Logger(server.UserTasks))
	server.router.POST("/user/:userid/task", Logger(server.TaskUpdate))
	server.router.DELETE("/user/:userid/task/:taskid", Logger(server.TaskDelete))
}

func main() {
	setup()
	routes()
	fmt.Println("Starting Server on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", server.router))
}

// Logger ...
func Logger(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		start := time.Now()

		handle(w, r, params)

		log.Printf(
			"%s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	}
}
