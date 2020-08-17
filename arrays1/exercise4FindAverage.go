package arrays1

import (
	"fmt"
	"os"
	"strconv"
)

// FindAverage testing
func FindAverage() {
	// ---------------------------------------------------------
	// EXERCISE: Find the Average
	//
	//  Your goal is to fill an array with numbers and find the average.
	//
	//  1. Get the numbers from the command-line.
	//
	//  2. Create an array and assign the given numbers to that array.
	//
	//  3. Print the given numbers and their average.
	//
	// RESTRICTION
	//   + Maximum 5 numbers can be provided
	//   + If one of the arguments are not a valid number, skip it
	//
	// EXPECTED OUTPUT
	//   go run main.go
	//     Please tell me numbers (maximum 5 numbers).
	//
	//   go run main.go 1 2 3 4 5 6
	//     Please tell me numbers (maximum 5 numbers).
	//
	//   go run main.go 1 2 3 4 5
	//     Your numbers: [1 2 3 4 5]
	//     Average: 3
	//
	//   go run main.go 1 a 2 b 3
	//     Your numbers: [1 0 2 0 3]
	//     Average: 2
	// ---------------------------------------------------------

	if len(os.Args) < 2 || len(os.Args) > 6 {
		fmt.Println(" Please tell me numbers (maximum 5 numbers). ")
		return
	}

	args := os.Args[1:]

	var numbs []float64

	// Appending the nums
	for _, num := range args {
		parNum, err := strconv.ParseFloat(num, 64)
		if err != nil {

			continue
		}
		numbs = append(numbs, parNum)
	}

	var generSum float64
	// Suming the numbers
	for _, num := range numbs {
		generSum += num
	}

	total := float64(len(numbs))
	var average = generSum / total

	fmt.Println("Your numbers: ", numbs)
	fmt.Println("Average:", average)

}
