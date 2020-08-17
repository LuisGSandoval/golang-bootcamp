package arrays1

import (
	"fmt"
	"os"
	"strconv"
)

// Sorter testing
func Sorter() {
	// ---------------------------------------------------------
	// EXERCISE: Number Sorter
	//
	//  Your goal is to sort the given numbers from the command-line.
	//
	//  1. Get the numbers from the command-line.
	//
	//  2. Create an array and assign the given numbers to that array.
	//
	//  3. Sort the given numbers and print them.
	//
	// RESTRICTION
	//   + Maximum 5 numbers can be provided
	//   + If one of the arguments are not a valid number, skip it
	//
	// HINTS
	//  + You can use the bubble-sort algorithm to sort the numbers.
	//    Please watch this: https://youtu.be/nmhjrI-aW5o?t=7
	//
	//  + When swapping the elements, do not check for the last element.
	//
	//    Or, you will receive this error:
	//    "panic: runtime error: index out of range"
	//
	// EXPECTED OUTPUT
	//   go run main.go
	//     Please give me up to 5 numbers.
	//
	//   go run main.go 6 5 4 3 2 1
	//     Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...
	//
	//   go run main.go 5 4 3 2 1
	//     [1 2 3 4 5]
	//
	//   go run main.go 5 4 a c 1
	//     [0 0 1 4 5]
	// ---------------------------------------------------------

	l := len(os.Args)
	if l < 2 {
		fmt.Println("Please give me up to 5 numbers.")
		return
	} else if l > 6 {
		fmt.Println("Sorry. Go arrays are fixed. So, for now, I'm only supporting sorting 5 numbers...")
		return
	}

	args := os.Args[1:]
	var validNums [5]float64

	for i, num := range args {
		validNum, err := strconv.ParseFloat(num, 64)
		if err != nil {
			validNums[i] = 0
			continue
		}
		validNums[i] = validNum
	}

	totalRounds := len(validNums)

keepSorting:
	for i := 0; i < len(validNums); i++ {
		if i == len(validNums)-1 {
			break
		}
		if validNums[i] > validNums[i+1] {
			validNums[i], validNums[i+1] = validNums[i+1], validNums[i]
		}
	}

	totalRounds--
	if totalRounds > 0 {
		goto keepSorting
	}

	fmt.Println(validNums)

}
