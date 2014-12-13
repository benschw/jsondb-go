package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

// Job to add a Todo to the database
type SaveTodoJob struct {
	toSave   Todo
	saved    chan Todo
	exitChan chan error
}

func NewSaveTodoJob(todo Todo) *SaveTodoJob {
	return &SaveTodoJob{
		toSave:   todo,
		saved:    make(chan Todo, 1),
		exitChan: make(chan error, 1),
	}
}
func (j SaveTodoJob) ExitChan() chan error {
	return j.exitChan
}
func (j SaveTodoJob) Run(todos map[string]Todo) (map[string]Todo, error) {
	var todo Todo
	if j.toSave.Id == "" {
		id, err := newUUID()
		if err != nil {
			return nil, err
		}
		todo = Todo{Id: id, Value: j.toSave.Value}
	} else {
		todo = j.toSave
	}
	todos[todo.Id] = todo

	j.saved <- todo
	return todos, nil
}

// Generate a uuid to use as a unique identifier for each Todo
// http://play.golang.org/p/4FkNSiUDMg
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
