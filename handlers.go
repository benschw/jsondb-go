package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type TodoHandlers struct {
	Client *DbClient
}

func (h *TodoHandlers) addTodo(res http.ResponseWriter, req *http.Request) {
	todo := &Todo{}
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(todo); err != nil {
		log.Print(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	created, err := h.Client.addTodo(*todo)
	if err != nil {
		log.Print(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(created)
	if err != nil {
		log.Print(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.Header().Set("Location", "todo/"+created.Id)
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, string(b[:]))
}

func (h *TodoHandlers) getTodos(res http.ResponseWriter, req *http.Request) {
	todos, err := h.Client.getTodos()
	if err != nil {
		log.Print(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(todos)
	if err != nil {
		log.Print(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, string(b[:]))
}

func (h *TodoHandlers) getTodo(res http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	todo, err := h.Client.getTodo(id)
	if err != nil {
		log.Print(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(todo)
	if err != nil {
		log.Print(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, string(b[:]))
}
