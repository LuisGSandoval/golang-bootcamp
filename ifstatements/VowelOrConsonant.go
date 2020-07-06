package ifstatements

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://en.oxforddictionaries.com/explore/is-the-letter-y-a-vowel-or-a-consonant/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

// VowelOrConsonant is a test function
func VowelOrConsonant() {

	const (
		consonants = "bcdfghjklmnpqrstvxz"
		vowels     = "aeiou"
		semiVowels = "wy"
	)

	fistArg := os.Args[1]

	// Error checking
	if len(fistArg) > 1 || len(os.Args) > 2 || len(os.Args) == 1 {
		fmt.Println("Give me a letter")
		return
	}

	// actuall feature
	consonant := strings.IndexAny(consonants, fistArg)
	vowel := strings.IndexAny(vowels, fistArg)
	semiVowel := strings.IndexAny(semiVowels, fistArg)

	if consonant > -1 {
		fmt.Printf("%s is a consonant \n", fistArg)
	} else if vowel > -1 {
		fmt.Printf("%s is a vowel \n", fistArg)
	} else if semiVowel > -1 {
		fmt.Printf("%s is a semivowel \n", fistArg)
	} else {
		fmt.Printf("%s is another character \n", fistArg)
	}

}
