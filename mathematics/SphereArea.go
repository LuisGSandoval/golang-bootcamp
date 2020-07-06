package mathematics

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// SphereArea calculates the area of a sphere by getting the radius
func SphereArea() {
	fmt.Println("Sphere area calculation")

	arg := os.Args[1]

	radius, _ := strconv.ParseFloat(arg, 64)

	area := 4 * math.Pi * (math.Pow(radius, 2))

	fmt.Printf("radius: %g -> area: %.2f\n", radius, area)

}
