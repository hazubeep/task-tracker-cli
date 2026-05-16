package main

func main() {

	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	Execute(&todos)
	storage.Save(todos)
}
