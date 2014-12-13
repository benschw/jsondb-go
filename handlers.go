package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type TodoHandlers struct {
	Client *TodoClient
}

// Add a new todo
func (h *TodoHandlers) AddTodo(c *gin.Context) {
	var todo Todo
	if !c.Bind(&todo) {
		c.JSON(400, "problem decoding body")
		return
	}
	todo.Id = ""

	created, err := h.Client.SaveTodo(todo)
	if err != nil {
		log.Print(err)
		c.JSON(500, "problem decoding body")
		return
	}

	c.JSON(201, created)
}

// Get all todos as an array
func (h *TodoHandlers) GetTodos(c *gin.Context) {
	todos, err := h.Client.GetTodos()
	if err != nil {
		log.Print(err)
		c.JSON(500, "problem decoding body")
		return
	}

	c.JSON(200, todos)
}

// Get a specific todo by id
func (h *TodoHandlers) GetTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	todo, err := h.Client.GetTodo(id)
	if err != nil {
		log.Print(err)
		c.JSON(500, "problem decoding body")
		return
	}

	c.JSON(200, todo)
}

// Add a new todo
func (h *TodoHandlers) SaveTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo Todo
	if !c.Bind(&todo) {
		c.JSON(400, "problem decoding body")
		return
	}
	todo.Id = id

	saved, err := h.Client.SaveTodo(todo)
	if err != nil {
		log.Print(err)
		c.JSON(500, "problem decoding body")
		return
	}

	c.JSON(200, saved)
}

// Delete a todo by id
func (h *TodoHandlers) DeleteTodo(c *gin.Context) {
	id := c.Params.ByName("id")

	if err := h.Client.DeleteTodo(id); err != nil {
		log.Print(err)
		c.JSON(500, "problem decoding body")
		return
	}

	c.Data(204, "application/json", make([]byte, 0))
}
