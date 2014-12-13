package main

import (
	"encoding/json"
	"io/ioutil"
)

// Job to delete a Todo from the database
type DeleteTodoJob struct {
	toDelete string
	exitChan chan error
}

func NewDeleteTodoJob(id string) *DeleteTodoJob {
	return &DeleteTodoJob{
		toDelete: id,
		exitChan: make(chan error, 1),
	}
}
func (j DeleteTodoJob) ExitChan() chan error {
	return j.exitChan
}
func (j DeleteTodoJob) Run(db string) error {
	todos := make(map[string]Todo, 0)
	content, err := ioutil.ReadFile(db)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(content, &todos); err != nil {
		return err
	}

	delete(todos, j.toDelete)

	b, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(db, b, 0644)
}
