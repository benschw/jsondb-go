package main

import (
	"encoding/json"
	"io/ioutil"
)

// Job to read all todos from the database
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
