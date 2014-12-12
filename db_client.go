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

	todos, err := c.getTodoHash()
	if err != nil {
		return arr, err
	}

	for _, value := range todos {
		arr = append(arr, value)
	}
	return arr, nil
}

func (c *DbClient) getTodo(id string) (Todo, error) {
	todos, err := c.getTodoHash()
	if err != nil {
		return Todo{}, err
	}
	return todos[id], nil
}

func (c *DbClient) getTodoHash() (map[string]Todo, error) {
	job := NewReadTodosJob()
	c.Jobs <- job

	if err := <-job.ExitChan(); err != nil {
		return make(map[string]Todo, 0), err
	}
	return <-job.todos, nil
}
