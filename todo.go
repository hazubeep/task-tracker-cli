package main

import (
	"errors"
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

	for index, todo := range *todos {
		if status == "" || todo.Status == status {
			table.AddRow(strconv.Itoa(index), todo.Description, string(todo.Status), todo.CreatedAt.Format(time.RFC1123), todo.UpdatedAt.Format(time.RFC1123))
		}
	}

	table.Render()
}

func (todos *Todos) add(description string) {
	now := time.Now()
	todo := Todo{
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	*todos = append(*todos, todo)

	// Todo: Change static ID
	fmt.Println("Task added successfully (ID: 1)")
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func (todos *Todos) edit(index int, newDescription string) error {

	if err := todos.validateIndex(index); err != nil {
		return err
	}

	(*todos)[index].Description = newDescription
	(*todos)[index].UpdatedAt = time.Now()

	return nil
}

func (todos *Todos) markInProgress(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	(*todos)[index].Status = StatusInProgress
	(*todos)[index].UpdatedAt = time.Now()

	return nil
}

func (todos *Todos) markDone(index int) error {

	if err := todos.validateIndex(index); err != nil {
		return err
	}

	(*todos)[index].Status = StatusDone
	(*todos)[index].UpdatedAt = time.Now()

	return nil
}
