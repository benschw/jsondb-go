package main

import (
	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	r.POST("/todo", handlers.AddTodo)
	r.GET("/todo", handlers.GetTodos)
	r.GET("/todo/:id", handlers.GetTodo)

	r.Run(":8080")

}
