package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Execute(todos *Todos) {

	if len(os.Args) < 2 {
		fmt.Println(`usage:
  todo-cli list <status>
  todo-cli add <description>
  todo-cli edit <id> <description>
  todo-cli delete <id>
  todo-cli mark-done <id>
  todo-cli mark-in-progress <id>
	`)
		return
	}

	command := os.Args[1]

	switch command {
	case "list":

		if len(os.Args) == 2 {

			todos.list(Status(""))
			return
		}

		status := os.Args[2]

		todos.list(Status(status))

	case "add":
		if len(os.Args) < 3 {
			fmt.Printf("Usage: app %s <id>\n", command)
			return
		}
		description := strings.Join(os.Args[2:], " ")
		todos.add(description)

	case "delete", "mark-in-progress", "mark-done":
		if len(os.Args) < 3 {
			fmt.Printf("Usage: app %s <id>", command)
			return
		}

		index, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Printf("Invalid ID")
			return
		}

		if command == "delete" {
			todos.delete(index)

		}
		if command == "mark-in-progress" {
			todos.markInProgress(index)

		}
		if command == "mark-done" {
			todos.markDone(index)

		}

	case "edit":
		if len(os.Args) < 4 {
			fmt.Printf("Usage: app %s <id> <description>", command)
			return
		}

		index, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Printf("Invalid ID")
			return
		}

		description := strings.Join(os.Args[3:], " ")

		todos.edit(index, description)

	default:
		fmt.Println("invalid command")
	}
}
