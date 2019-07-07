package todo

import "fmt"

type Todo struct {
	ID string `json:"id"`
	Task string `task:"id"`
}

var Todos = []*Todo{
	{ ID: "todo1", Task: "build an API"},
	{ ID: "todo2", Task: "?????"},
	{ ID: "todo3", Task: "profit!"},
}

func GetTodo(todoID string) (*Todo, error) {

	for _, todo := range Todos {
		if todo.ID == todoID {
			return todo, nil
		}
	}
	return nil, fmt.Errorf("Todo not found: %s", todoID)
}

func GetTodos() []*Todo {
	return Todos
}
