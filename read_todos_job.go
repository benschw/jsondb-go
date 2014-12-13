package main

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
func (j ReadTodosJob) Run(todos map[string]Todo) (map[string]Todo, error) {
	j.todos <- todos

	return nil, nil
}
