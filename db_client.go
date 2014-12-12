package main

type DbClient struct {
	Jobs chan Job
}

func (c *DbClient) addTodo(todo Todo) (Todo, error) {
	job := NewAddTodoJob(todo)
	c.Jobs <- job

	if err := <-job.ExitChan(); err != nil {
		return Todo{}, nil
	}
	return <-job.created, nil
}

func (c *DbClient) getTodos() ([]Todo, error) {
	arr := make([]Todo, 0)

	job := NewReadTodosJob()
	c.Jobs <- job

	if err := <-job.ExitChan(); err != nil {
		return arr, err
	}
	todos := <-job.todos

	for _, value := range todos {
		arr = append(arr, value)
	}
	return arr, nil
}
