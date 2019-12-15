package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", UsersList)
	router.POST("/user", UserUpdate)
	router.GET("/user/:userid", UserTasks)
	router.DELETE("/user/:userid", UserDelete)

	router.POST("/user/:userid/task", TaskUpdate)
	router.DELETE("/user/:userid/task/:taskid", TaskDelete)

	fmt.Println("Starting Server on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
