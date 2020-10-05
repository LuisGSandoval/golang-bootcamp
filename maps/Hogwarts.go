package maps

import (
	"fmt"
	"os"
	"sort"
)

// Hogwarts testing
func Hogwarts() {

	hogwarts := map[string][]string{
		"gryffindor": {"weasley", "hagrid", "dumbledore", "lupin"},
		"hufflepuf":  {"wenlock", "scamander", "helga", "diggory"},
		"ravenclaw":  {"flitwick", "bagnold", "wildsmith", "montmorency"},
		"slytherin":  {"horace", "nigellus", "higgs", "scorpius"},
		"bobo":       {"wizardry", "unwanted"},
	}

	// Checking valid information

	delete(hogwarts, "bobo")

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Please select one hogwarts house name [gryffindor,hufflepuf, ravenclaw, slytherin]")
		return
	}

	students, house := hogwarts[args[0]], args[0]

	if students == nil {

		fmt.Printf("%s is not part of wogwarts\n", house)
		return
	}

	fmt.Printf("~~~ %s's students ~~~\n\n", house)

	sort.Strings(students)
	for _, val := range students {
		fmt.Println(val)
	}

}
