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

}

func routes() {
	server.router.GET("/users", Logger(server.UsersList))
	server.router.OPTIONS("/user", Logger(server.Options))
	server.router.POST("/user", Logger(server.UserUpdate))
	server.router.OPTIONS("/user/:userid", Logger(server.Options))
	server.router.DELETE("/user/:userid", Logger(server.UserDelete))

	server.router.GET("/tasks/:userid", Logger(server.UserTasks))
	server.router.OPTIONS("/task/:userid", Logger(server.Options))
	server.router.POST("/task/:userid", Logger(server.TaskUpdate))
	server.router.OPTIONS("/task/:userid/:taskid", Logger(server.Options))
	server.router.DELETE("/task/:userid/:taskid", Logger(server.TaskDelete))
}

func main() {
	setup()
	routes()
	fmt.Println("Starting Server on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", server.router))
}

// Options ...
func (server *Server) Options(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	CORSHeaders(&w, r)
}

// CORSHeaders ...
func CORSHeaders(w *http.ResponseWriter, r *http.Request) {
	// For test purposes only,
	// Set more fine grained locking in production
	origin := r.Header.Get("Origin")
	if origin == "" {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
	} else {
		(*w).Header().Set("Access-Control-Allow-Origin", origin)
	}
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")
}

// Logger ...
func Logger(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		CORSHeaders(&w, r)
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
