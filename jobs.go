package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

type Job interface {
	ExitChan() chan error
	Run(db string) error
}

// Job to write the database
type AddJob struct {
	todo     chan Todo
	created  chan Todo
	exitChan chan error
}

func NewAddJob(todo Todo) *AddJob {
	j := &AddJob{
		todo:     make(chan Todo, 1),
		created:  make(chan Todo, 1),
		exitChan: make(chan error, 1),
	}
	j.todo <- todo
	return j
}
func (j AddJob) ExitChan() chan error {
	return j.exitChan
}
func (j AddJob) Run(db string) error {
	todos := make(map[string]Todo, 0)
	content, err := ioutil.ReadFile(db)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(content, &todos); err != nil {
		return err
	}
	todo := <-j.todo
	id, err := newUUID()
	todo.Id = id
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

// Job to read the database
type ReadTodosJob struct {
	todos    chan map[string]Todo
	exitChan chan error
}

func NewReadTodosJob() *ReadTodosJob {
	return &ReadTodosJob{
		todos:    make(chan map[string]Todo, 1),
		exitChan: make(chan error, 1),
	}
}
func (j ReadTodosJob) ExitChan() chan error {
	return j.exitChan
}
func (j ReadTodosJob) Run(db string) error {
	todos := make(map[string]Todo, 0)
	content, err := ioutil.ReadFile(db)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(content, &todos); err != nil {
		return err
	}

	j.todos <- todos

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
