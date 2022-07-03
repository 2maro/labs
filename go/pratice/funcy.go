package main

import (
	"fmt"
	"math"
)

func main() {
	x := sum(5, 6)

	b := 9.0
	c, err := sqrt(b)
	if err != nil {
		fmt.Println("There was error")
	}
	fmt.Println(c)

	fmt.Println("sum:", x)
}

func sum(y, x int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("x cannot be less then zero")
	}
	return math.Sqrt(x), nil
}
