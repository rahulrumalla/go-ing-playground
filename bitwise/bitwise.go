package main

import (
	"fmt"
	"math"
)

func main() {
	ff := fmt.Printf

	var x float64 = 2
	x = x * math.Pow(2, 4)
	ff("%v\n", x)

	y := 2 << 4
	ff("%v\n", y)
}
