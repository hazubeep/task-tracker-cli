package main

import "fmt"

func main() {

	todos := Todos{}

	todos.add("Buy Toys")
	todos.add("Buy Breed")
	fmt.Printf("%+v\n\n", todos)

	todos.edit(1, "Buy Milk")
	fmt.Printf("%+v\n\n", todos)

	todos.delete(0)
	fmt.Printf("%+v\n", todos)

}
