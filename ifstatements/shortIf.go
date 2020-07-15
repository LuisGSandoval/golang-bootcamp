package ifstatements

import (
	"fmt"
	"os"
	"strconv"
)

// ShortIf is a test function
func ShortIf() {

	var (
		n   int
		err error
	)

	if a := os.Args; len(a) != 2 {
		fmt.Println("Give me a number")
	} else if n, err = strconv.Atoi(a[1]); err != nil {
		fmt.Printf("Can't convert %q. \n", a[1])
	} else if n == 32 {
		fmt.Printf("Perfect number, %v , %v \n", a, n)
	}

	fmt.Println(n)

}
