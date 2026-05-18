# Task Tracker CLI

A command-line interface (CLI) application to track and manage your tasks, built with Go. This is a solution for the [Task Tracker](https://roadmap.sh/projects/task-tracker) project on roadmap.sh.

## Features

- Add, update, and delete tasks
- Mark tasks as in-progress or done
- List all tasks or filter by status
- Persistent storage using a local `todo.json` file
- Graceful error and edge case handling

## Requirements

- [Go](https://golang.org/dl/) 1.21 or higher
- [github.com/aquasecurity/table](https://github.com/aquasecurity/table) — for table rendering

## Installation

**1. Clone the repository**

```bash
git clone https://github.com/hazubeep/task-tracker-cli.git
cd task-tracker-cli
```

**2. Install dependencies**

```bash
go mod tidy
```

**3. Build the binary**

```bash
go build -o task-cli .
```

Or run directly without building:

```bash
go run main.go <command>
```

## Usage

### Add a Task

```bash
./task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

### Update a Task

```bash
./task-cli update 1 "Buy groceries and cook dinner"
```

### Delete a Task

```bash
./task-cli delete 1
```

### Mark Task as In Progress

```bash
./task-cli mark-in-progress 1
```

### Mark Task as Done

```bash
./task-cli mark-done 1
```

### List Tasks

```bash
# List all tasks
./task-cli list

# List only completed tasks
./task-cli list done

# List only pending tasks
./task-cli list todo

# List tasks in progress
./task-cli list in-progress
```

## Task Properties

Each task stored in `todo.json` contains the following fields:

| Field         | Type   | Description                              |
|---------------|--------|------------------------------------------|
| `id`          | int    | Unique identifier (auto-incremented)     |
| `description` | string | Short description of the task            |
| `status`      | string | `todo`, `in_progress`, or `done`         |
| `createdAt`   | time   | Timestamp when the task was created      |
| `updatedAt`   | time   | Timestamp when the task was last updated |

### Example `todo.json`

```json
[
  {
    "ID": 1,
    "Description": "Buy groceries",
    "Status": "done",
    "CreatedAt": "2024-01-01T10:00:00Z",
    "UpdatedAt": "2024-01-01T12:00:00Z"
  },
  {
    "ID": 2,
    "Description": "Cook dinner",
    "Status": "in_progress",
    "CreatedAt": "2024-01-01T11:00:00Z",
    "UpdatedAt": "2024-01-01T11:30:00Z"
  }
]
```

## Project Structure

```
task-tracker-cli/
├── main.go        # Entry point
├── todo.go        # Todo struct, methods, and business logic
├── storage.go     # JSON read/write operations
├── todo.json      # Auto-generated task storage file
├── command.go     # CLI argument handling
├── go.mod
└── go.sum
```

## Error Handling

The application handles the following edge cases:

- Invalid or out-of-range task ID
- Missing or malformed `todo.json` file
- Incorrect number of arguments
- Unknown commands

## License

This project is open source and available under the [MIT License](LICENSE).
