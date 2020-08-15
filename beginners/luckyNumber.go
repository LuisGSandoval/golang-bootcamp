package beginners

import (
	"fmt"
	"math/rand"
	"time"
)

// GuessNumber :  testing functionality
func GuessNumber() {

	num := rand.Intn(5)
	for i := 0; i <= num; i++ {
		rand.Seed(time.Now().UnixNano())
		fmt.Println(num)
	}
}
