package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type TodoHandlers struct {
	Jobs chan Job
}

func (h *TodoHandlers) addTodo(res http.ResponseWriter, req *http.Request) {
	todo := &Todo{}
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(todo); err != nil {
		log.Print(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	job := NewAddJob(*todo)
	h.Jobs <- job

	if err := <-job.ExitChan(); err != nil {
		log.Print(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	created := <-job.created
	b, err := json.Marshal(created)
	if err != nil {
		log.Print(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Location", "todo/"+created.Id)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, string(b[:]))
}

func (h *TodoHandlers) getTodos(res http.ResponseWriter, req *http.Request) {
	job := NewReadTodosJob()
	h.Jobs <- job

	if err := <-job.ExitChan(); err != nil {
		log.Print(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	todos := <-job.todos

	arr := make([]Todo, 0)
	for _, value := range todos {
		arr = append(arr, value)
	}

	b, err := json.Marshal(arr)
	if err != nil {
		log.Print(err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, string(b[:]))
}
