package scaning

import (
	"bufio"
	"fmt"
	"os"
)

// Scanning1 testing
func Scanning1() {

	defer os.Stdin.Close()

	in := bufio.NewScanner(os.Stdin)

	var lines int
	for in.Scan() {

		lines++
		// fmt.Println("Texto:", in.Text())
	}
	fmt.Println("lines: ", lines)

	err := in.Err()

	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Println("Hex:", in.Bytes())
}
