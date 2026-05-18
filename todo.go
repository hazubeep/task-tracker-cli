package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Status string

const (
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

type Todo struct {
	ID          int
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Todos []Todo

func (todos *Todos) list(status Status) {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("id", "Description", "Status", "Created At", "Updated At")

	for _, todo := range *todos {
		if status == "" || todo.Status == status {
			table.AddRow(strconv.Itoa(todo.ID), todo.Description, string(todo.Status), todo.CreatedAt.Format(time.RFC1123), todo.UpdatedAt.Format(time.RFC1123))
		}
	}

	table.Render()
}

func (todos *Todos) add(description string) {
	now := time.Now()
	todo := Todo{
		ID:          todos.getNextId(),
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	*todos = append(*todos, todo)

	// Todo: Change static ID
	fmt.Printf("Task added successfully (ID: %s)", strconv.Itoa(todos.getNextId()))
}

func (todos *Todos) delete(id int) error {

	index, err := todos.findById(id)

	if err != nil {
		return err
	}

	t := *todos
	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func (todos *Todos) edit(id int, newDescription string) error {

	index, err := todos.findById(id)

	if err != nil {
		return err
	}

	(*todos)[index].Description = newDescription
	(*todos)[index].UpdatedAt = time.Now()

	return nil
}

func (todos *Todos) markInProgress(id int) error {

	index, err := todos.findById(id)

	if err != nil {
		return err
	}

	(*todos)[index].Status = StatusInProgress
	(*todos)[index].UpdatedAt = time.Now()

	return nil
}

func (todos *Todos) markDone(id int) error {

	index, err := todos.findById(id)

	if err != nil {
		return err
	}

	(*todos)[index].Status = StatusDone
	(*todos)[index].UpdatedAt = time.Now()

	return nil
}

func (todos *Todos) getNextId() int {
	maxID := 0

	for _, todo := range *todos {
		if todo.ID > maxID {
			maxID = todo.ID
		}
	}

	return maxID + 1
}

func (todos *Todos) findById(id int) (int, error) {
	for i, todo := range *todos {
		if todo.ID == id {
			return i, nil
		}
	}
	return -1, fmt.Errorf("ID not found")
}
