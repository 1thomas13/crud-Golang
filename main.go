package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", editTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))
}
