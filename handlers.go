package main

import (
	"github.com/gin-gonic/gin"
)

type Error struct {
	Error string `json:"error"`
}

type TodoHandlers struct {
	Client *DbClient
}

func (h *TodoHandlers) AddTodo(c *gin.Context) {
	var todo Todo
	if !c.Bind(&todo) {
		c.JSON(400, &Error{"problem decoding body"})
		return
	}

	created, err := h.Client.AddTodo(todo)
	if err != nil {
		c.JSON(500, &Error{"problem decoding body"})
		return
	}

	c.JSON(201, created)
}

func (h *TodoHandlers) GetTodos(c *gin.Context) {
	todos, err := h.Client.GetTodos()
	if err != nil {
		c.JSON(500, &Error{"problem decoding body"})
		return
	}

	c.JSON(200, todos)
}

func (h *TodoHandlers) GetTodo(c *gin.Context) {
	id := c.Params.ByName("id")
	todo, err := h.Client.GetTodo(id)
	if err != nil {
		c.JSON(500, &Error{"problem decoding body"})
		return
	}

	c.JSON(200, todo)
}
