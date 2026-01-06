package main

import "fmt"

func main() {
	todos := Todos{}
	todos.add("get this")
	todos.add("get that")
	fmt.Printf("%+v\n\n", todos)
	todos.delete(0)
	fmt.Printf("%+v\n\n", todos)

}
