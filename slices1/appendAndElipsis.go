package slices1

import (
	"github.com/inancgumus/prettyslice"
)

// AppendAndElipsis testing
func AppendAndElipsis() {
	var todo []string

	var otrosTodos = []string{
		"play ukulele",
		"Drik soda",
		"Sleep like a log",
	}

	todo = append(todo, otrosTodos...)
	todo = append(todo, "cantar", "comer pizza")

	prettyslice.Show("Hola Luis, estos son tus tareas", todo)

}
