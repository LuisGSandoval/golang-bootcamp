package beginners

import (
	"fmt"
	"os"
	"strings"
)

// WordFinder : testing
func WordFinder() {
	const corpus = "LAZY cat jumps over and over again and again"

	words := strings.Fields(corpus)
	query := os.Args[1:]

	// this is done to find words entered in the preadded string
	for _, q := range query {
		for i, word := range words {
			if strings.ToLower(q) == strings.ToLower(word) {
				fmt.Printf("#%d  %q \n", i+1, word)
				break
			}
		}
	}

}
