package stringrunebyte

import (
	"fmt"
	"strings"
)

// Stringrunebyte testing
func Stringrunebyte() {
	beginning, end := 33, 255

	fmt.Printf("%-6s => %-6s => %-6s => %-6s => %-6s\n %s \n", "representation", "Lit", "Dec", "Hex", "Encoded", strings.Repeat("-", 45))

	for i := beginning; i <= end; i++ {

		fmt.Printf("%-4s => %-6[1]c => %-6[1]d => %-6[1]x => % -12x \n", i, string(i))

	}
}
