package main

import "fmt"

type task struct {
	name        string
	description string
	completed   bool
}

func (t task) setCompleted() {
	t.completed = true
}

func (t task) setName(name string) {
	t.name = name
}

func (t task) setDescription(description string) {
	t.description = description
}

func main() {
	t := task{
		name:        "Completar curso Go",
		description: "Url aqui",
		completed:   false,
	}

	fmt.Println(t)
}
