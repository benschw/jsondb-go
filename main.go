package main

import (
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

const Db = "./db.json"

func main() {

	if _, err := ioutil.ReadFile(Db); err != nil {
		str := "{}"
		if err = ioutil.WriteFile(Db, []byte(str), 0644); err != nil {
			log.Fatal(err)
		}
	}

	jobs := make(chan Job)
	go ProcessJobs(jobs, Db)

	client := &DbClient{Jobs: jobs}
	handlers := &TodoHandlers{Client: client}

	r := mux.NewRouter()

	r.HandleFunc("/todo", handlers.addTodo).Methods("POST")
	r.HandleFunc("/todo", handlers.getTodos).Methods("GET")
	r.HandleFunc("/todo/{id}", handlers.getTodo).Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
