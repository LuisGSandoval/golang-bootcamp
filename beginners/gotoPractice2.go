package beginners

import "fmt"

// GotoPractice2 : testing
func GotoPractice2() {

	goto exit

start:
	goto getout

exit:
	goto start

getout:
	fmt.Println("exiting")

}
