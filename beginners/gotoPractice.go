package beginners

import "fmt"

// GotoPractice : testing
func GotoPractice() {

	var i int

looping:

	if i < 3 {
		fmt.Println("Looping: ", i)
		i++
		goto looping
	}

	fmt.Println("finished")

}
