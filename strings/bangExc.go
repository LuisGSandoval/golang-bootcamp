package strings

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

// Echo is a gfunction that returns a string as a scream!!
func Echo() {
	arg := os.Args[1]

	l := utf8.RuneCountInString(arg)

	exc := strings.Repeat("!", l)

	s := exc + strings.ToUpper(arg) + exc

	fmt.Println(s)
}
