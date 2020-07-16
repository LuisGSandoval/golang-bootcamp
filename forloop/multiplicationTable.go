package forloop

import "fmt"

// MultiplicationTable is ....
func MultiplicationTable() {

	max := 5

	fmt.Printf("%5s", "x")

	for i := 0; i < max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", j*i)
		}

		fmt.Println()
	}

}
