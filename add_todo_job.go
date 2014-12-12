package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

// Job to add a Todo to the database
type AddTodoJob struct {
	toAdd    Todo
	created  chan Todo
	exitChan chan error
}

func NewAddTodoJob(todo Todo) *AddTodoJob {
	return &AddTodoJob{
		toAdd:    todo,
		created:  make(chan Todo, 1),
		exitChan: make(chan error, 1),
	}
}
func (j AddTodoJob) ExitChan() chan error {
	return j.exitChan
}
func (j AddTodoJob) Run(db string) error {
	todos := make(map[string]Todo, 0)
	content, err := ioutil.ReadFile(db)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(content, &todos); err != nil {
		return err
	}

	id, err := newUUID()
	todo := Todo{Id: id, Value: j.toAdd.Value}
	todos[id] = todo

	b, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(db, b, 0644); err != nil {
		return err
	}

	j.created <- todo
	return nil
}

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
