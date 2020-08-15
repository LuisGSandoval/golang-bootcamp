package beginners

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Crunch the primes
//
//  1. Get numbers from the command-line.
//  2. `for range` over them.
//  4. Convert each one to an int.
//  5. If one of the numbers isn't an int, just skip it.
//  6. Print the ones that are only the prime numbers.
//
// RESTRICTION
//  The user can run the program with any number of
//  arguments.
//
// HINT
//  Find whether a number is prime using this algorithm:
//  https://stackoverflow.com/a/1801446/115363
//
// EXPECTED OUTPUT
//  go run main.go 2 4 6 7 a 9 c 11 x 12 13
//    2 7 11 13
//
//  go run main.go 1 2 3 5 7 A B C
//    2 3 5 7
// ---------------------------------------------------------

// CrunchThePrimes testing
func CrunchThePrimes() {

	nums := os.Args[1:]

	var primeNums []int

	for _, num := range nums {
		posiblePrime, err := strconv.Atoi(num)
		if err != nil {
			continue
		}

		if isPrime(posiblePrime) {
			primeNums = append(primeNums, posiblePrime)
		}
	}

	fmt.Println(primeNums)

}

func isPrime(n int) bool {

	switch {
	case n < 2:
		return false
	case 2 == n, 3 == n:
		return true
	case n%3 == 0, n%2 == 0:
		return false
	}

	for i := 5; i*i <= n; i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}
