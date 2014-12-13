package main

type TodoClient struct {
	Jobs chan Job
}

func (c *TodoClient) SaveTodo(todo Todo) (Todo, error) {
	job := NewSaveTodoJob(todo)
	c.Jobs <- job

	if err := <-job.ExitChan(); err != nil {
		return Todo{}, err
	}
	return <-job.saved, nil
}

func (c *TodoClient) GetTodos() ([]Todo, error) {
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

func (c *TodoClient) GetTodo(id string) (Todo, error) {
	todos, err := c.getTodoHash()
	if err != nil {
		return Todo{}, err
	}
	return todos[id], nil
}

func (c *TodoClient) DeleteTodo(id string) error {
	job := NewDeleteTodoJob(id)
	c.Jobs <- job

	if err := <-job.ExitChan(); err != nil {
		return err
	}
	return nil
}

func (c *TodoClient) getTodoHash() (map[string]Todo, error) {
	job := NewReadTodosJob()
	c.Jobs <- job

	if err := <-job.ExitChan(); err != nil {
		return make(map[string]Todo, 0), err
	}
	return <-job.todos, nil
}
