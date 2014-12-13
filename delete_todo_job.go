package main

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
func (j DeleteTodoJob) Run(todos map[string]Todo) (map[string]Todo, error) {
	delete(todos, j.toDelete)
	return todos, nil
}
