package main

import "fmt"

type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) setCompleted() {
	t.completed = true
}

func (t *task) setName(name string) {
	t.name = name
}

func (t *task) setDescription(description string) {
	t.description = description
}

type taskList struct {
	tasks []*task
}

func (tl *taskList) addToList(t *task) {
	tl.tasks = append(tl.tasks, t)
}

func (tl *taskList) removeToList(index int) {
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

func main() {

	lista := &taskList{}

	t1 := &task{
		name:        "Completar curso Go",
		description: "Hacerlo esta semana",
		completed:   false,
	}

	t2 := &task{
		name:        "Completar curso Ingles",
		description: "Hacerlo este mes",
		completed:   false,
	}

	t3 := &task{
		name:        "Completar curso K8s",
		description: "Hacerlo esta año",
		completed:   false,
	}

	lista.addToList(t1)
	lista.addToList(t2)
	lista.addToList(t3)

	fmt.Println(lista.tasks[0])

}
