package emptyfilefinder

import (
	"fmt"
	"io/ioutil"
	"os"
)

// EmptyFileFinder testing
func EmptyFileFinder() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println(` Please add a directory `)
		return
	}

	files, err := ioutil.ReadDir(args[0])

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("The emtpy files are")
	fmt.Println("====================")

	// total := len(files) * (255 + 1)

	var total int

	for _, file := range files {
		if file.Size() == 0 {
			total += len(file.Name())
		}
	}

	names := make([]byte, 0, total)

	fmt.Printf("size %d bytes \n", total)

	for _, file := range files {
		if file.Size() == 0 {
			names = append(names, file.Name()...)
			names = append(names, "\n"...)
		}
	}

	err = ioutil.WriteFile("out.txt", names, 0644)

	if err != nil {
		fmt.Println("Error when writing to out.txt file")
		return
	}

	fmt.Printf("%s", names)
}
