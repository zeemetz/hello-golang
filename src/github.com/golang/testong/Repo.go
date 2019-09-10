package main

import "fmt"

var currentID int
var todos Todos

func init() {
	RepoCreateTodo(Todo{Name: "Testong", Due: "2017-01-01 00:00:00"})
	RepoCreateTodo(Todo{Name: "Host Meetup", Due: "2018-01-01 00:00:00"})
}

// RepoCreateTodo is mean to do...
func RepoCreateTodo(t Todo) Todo {
	currentID++
	t.ID = currentID
	todos = append(todos, t)

	//save todo to database
	return t
}

// RepoFindTodo is meant to do
func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}

	return Todo{}
}

// RepoDestroyTodo is mean to do
func RepoDestroyTodo(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
