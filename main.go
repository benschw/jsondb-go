package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const Db = "./db.json"

func main() {

	jobs := make(chan Job)
	go ProcessJobs(jobs, Db)

	handlers := &TodoHandlers{Jobs: jobs}

	r := mux.NewRouter()

	r.HandleFunc("/todo", handlers.addTodo).Methods("POST")
	r.HandleFunc("/todo", handlers.getTodos).Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
