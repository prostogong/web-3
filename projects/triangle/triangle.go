package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Scan(&a)
	fmt.Scan(&b)

	c := math.Sqrt(a*a + b*b)

	fmt.Print(c)
}
