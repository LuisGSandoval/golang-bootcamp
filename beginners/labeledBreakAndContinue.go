package beginners

import (
	"fmt"
	"os"
	"strings"
)

// LabeledBreakAndContinue : testing
func LabeledBreakAndContinue() {
	const corpus = "LAZY cat jumps over and over again and again"

	words := strings.Fields(corpus)
	query := os.Args[1:]

queries: // this is done to find words entered in the preadded string
	for _, q := range query {

	search:
		for i, word := range words {

			switch q {
			case "or", "and":
				break search
			}

			if strings.ToLower(q) == strings.ToLower(word) {
				fmt.Printf("#%d  %q \n", i+1, word)
				break queries
			}
		}
	}

	// 	fmt.Print("\n\n\n*****************************\n\n\n")

	// queries2: // this is done to find words entered in the preadded string
	// 	for _, q := range query {
	// 		for i, word := range words {
	// 			if strings.ToLower(q) == strings.ToLower(word) {
	// 				fmt.Printf("#%d  %q \n", i+1, word)
	// 				continue queries2
	// 			}
	// 		}
	// 	}

}
