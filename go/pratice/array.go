package main

import "fmt"

func main() {
	sample := []int{1, 2, 3, 4, 5, 6}
	sample[1] = 5
	sample = append(sample, 10)
	fmt.Println(sample)
}
