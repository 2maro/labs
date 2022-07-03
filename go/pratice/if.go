package main

import "fmt"

func main() {
	blah := map[string]int{
		"blaa": 1,
		"do":   3,
		"This": 2,
	}
	if blah["do"] < 2 {
		fmt.Println("greater than 3")
	} else if blah["do"] > 3 {
		fmt.Println(" not greater then 3")
	} else {
		fmt.Println("else")
	}
}
