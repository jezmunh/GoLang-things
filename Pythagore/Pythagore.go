package main

import (
	"fmt"
	"math"
)

func find_hypotenuse(leg1 float64, leg2 float64) int {
	return int(math.Sqrt(math.Pow(leg1, 2) + math.Pow(leg2, 2)))
}

func main() {

	fmt.Println("The hypotenuse of the triangle:", find_hypotenuse(3, 4))
}
