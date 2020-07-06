package mathematics

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// ShepreVolume is a function that claculates the volume of a sphere with a given radius
func ShepreVolume() {

	arg := os.Args[1]

	radius, _ := strconv.ParseFloat(arg, 64)

	vol := (4 * math.Pi * math.Pow(radius, 3)) / 3

	fmt.Printf("radius: %g -> volume: %g\n", radius, vol)

}
