package emptyfilefinder

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

// WordSorterFileWriter testing
func WordSorterFileWriter() {
	args := os.Args[1:]

	sort.Strings(args)

	var words []byte

	for i, arg := range args {
		words = append(words, strconv.Itoa(i+1)+". "...)
		words = append(words, arg...)
		words = append(words, "\n"...)
	}

	err := ioutil.WriteFile("sortedArgs.txt", words, 0644)

	if err != nil {
		fmt.Println("Error when writing to sortedArgs.txt file")
		return
	}

}
