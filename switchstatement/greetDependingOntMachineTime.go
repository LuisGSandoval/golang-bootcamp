package switchstatement

import (
	"fmt"
	"time"
)

// SayHi this is a test function
func SayHi() {

	switch rn := time.Now().Hour(); {
	case rn > 18:
		fmt.Println("Good evening")
	case rn > 12:
		fmt.Println("Good afternoon")
	case rn > 0:
		fmt.Println("Good morning")

	default:
		fmt.Println("Hello")
	}

}
