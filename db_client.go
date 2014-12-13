package main

type DbClient struct {
	Jobs chan Job
}

func (c *DbClient) AddTodo(todo Todo) (Todo, error) {
	job := NewAddTodoJob(todo)
	c.Jobs <- job

	if err := <-job.ExitChan(); err != nil {
		return Todo{}, err
	}
	return <-job.created, nil
}

func (c *DbClient) GetTodos() ([]Todo, error) {
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

func (c *DbClient) GetTodo(id string) (Todo, error) {
	todos, err := c.getTodoHash()
	if err != nil {
		return Todo{}, err
	}
	return todos[id], nil
}

func (c *DbClient) DeleteTodo(id string) error {
	job := NewDeleteTodoJob(id)
	c.Jobs <- job

	if err := <-job.ExitChan(); err != nil {
		return err
	}
	return nil
}

func (c *DbClient) getTodoHash() (map[string]Todo, error) {
	job := NewReadTodosJob()
	c.Jobs <- job

	if err := <-job.ExitChan(); err != nil {
		return make(map[string]Todo, 0), err
	}
	return <-job.todos, nil
}
