package main

import "fmt"

type filerType string

const (
	All       filerType = "all"
	Pending   filerType = "pending"
	Completed filerType = "completed"
)

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

func (tl *taskList) printTasks(filter filerType) {
	for i, task := range tl.tasks {
		if filter == Completed {
			if task.completed {
				fmt.Println("Index:", i)
				fmt.Println("Nombre:", task.name)
				fmt.Println("Descripción:", task.description)
				fmt.Println("Completa?:", task.completed)
				fmt.Println("")
			}
		}

		if filter == Pending {
			if !task.completed {
				fmt.Println("Index:", i)
				fmt.Println("Nombre:", task.name)
				fmt.Println("Descripción:", task.description)
				fmt.Println("Completa?:", task.completed)
				fmt.Println("")
			}
		}

		if filter == All {
			fmt.Println("Index:", i)
			fmt.Println("Nombre:", task.name)
			fmt.Println("Descripción:", task.description)
			fmt.Println("Completa?:", task.completed)
			fmt.Println("")
		}

	}
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
		completed:   true,
	}

	t3 := &task{
		name:        "Completar curso K8s",
		description: "Hacerlo esta año",
		completed:   false,
	}

	lista.addToList(t1)
	lista.addToList(t2)
	lista.addToList(t3)

	lista.printTasks(Pending)

}
