package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

const Db = "./db.json"

func main() {

	// initialize empty-object json file if not found
	if _, err := ioutil.ReadFile(Db); err != nil {
		str := "{}"
		if err = ioutil.WriteFile(Db, []byte(str), 0644); err != nil {
			log.Fatal(err)
		}
	}

	// create channel to communicate over
	jobs := make(chan Job)

	// start watching jobs channel for work
	go ProcessJobs(jobs, Db)

	// create dependencies
	client := &DbClient{Jobs: jobs}
	handlers := &TodoHandlers{Client: client}

	// configure routes
	r := gin.Default()

	r.POST("/todo", handlers.AddTodo)
	r.GET("/todo", handlers.GetTodos)
	r.GET("/todo/:id", handlers.GetTodo)
	r.DELETE("/todo/:id", handlers.DeleteTodo)

	// start web server
	r.Run(":8080")
}
