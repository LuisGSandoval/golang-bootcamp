package mathematics

import (
	"fmt"
	"os"
	"strconv"
)

// CelsiusToFahrenheit is a function that converts celsius to fahrenheit
func CelsiusToFahrenheit() {

	fmt.Println("Temperature conversion")
	fmt.Println("Celsius to fahrenheit")

	arg := os.Args[1]

	celsius, _ := strconv.ParseFloat(arg, 64)

	fahrenheit := celsius*1.8 + 32

	fmt.Printf("%.1f ºC = %.1ff\n", celsius, fahrenheit)

}
