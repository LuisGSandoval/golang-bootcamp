package arrays1

import (
	"fmt"
	"os"
	"strings"
)

// WordFinder testing
func WordFinder() {

	// ---------------------------------------------------------
	// EXERCISE: Word Finder
	//
	//   Your goal is to search for the words inside the corpus.
	//
	//   Note: This exercise is similar to the previous word finder program:
	//   https://github.com/inancgumus/learngo/tree/master/13-loops/10-word-finder-labeled-switch
	//
	//   1. Get the search query from the command-line (it can be multiple words)
	//
	//   2. Filter these words, do not search for them:
	//      and, or, was, the, since, very
	//
	//      To do this, use an array for the filtered words.
	//
	//   3. Print the words found.
	//
	// RESTRICTION
	//   + The search and the filtering should be case insensitive
	//
	// HINT
	//   + strings.Fields function converts a given string to a slice.
	//
	//     You can find its example in the word finder program that I've mentioned
	//     above.
	//
	// EXPECTED OUTPUT
	//   go run main.go
	//     Please give me a word to search.
	//
	//   go run main.go and was
	//
	//   go run main.go AND WAS
	//
	//   go run main.go cat beginning
	//     #2 : "cat"
	//     #11: "beginning"
	//
	//   go run main.go Cat Beginning
	//     #2 : "cat"
	//     #11: "beginning"
	// ---------------------------------------------------------

	const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

	args := os.Args[1:]

	if l := len(args); l == 0 {
		fmt.Println("Please give me a word to search.")
	}

	wordsArr := strings.Fields(corpus)

	for _, arg := range args {

		switch strings.ToLower(arg) {
		case "and", "or", "was", "the", "since", "very":
			continue
		}

		for i, word := range wordsArr {
			if strings.ToLower(word) == strings.ToLower(arg) {
				fmt.Printf("#%-2d: %q \n", i+1, word)
			}
		}

	}

}
