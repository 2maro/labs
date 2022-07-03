package main

import "fmt"

func main() {
	//normal loop
	for i := 0; i < 5; i++ {
		//	fmt.Println(i)
	}

	blah := []int{1, 2, 3, 4, 5}
	for index, value := range blah {
		fmt.Println("index", index, "value", value)
	}

	//ignore index
	foo := []int{6, 7, 8, 9, 10}
	for _, value := range foo {
		fmt.Println("value", value)
	}

	//key:value storage
	bar := map[string]int{"one": 1, "two": 2}
	for key, value := range bar {
		fmt.Println("Key", key, "Value", value)
	}
	//infi-loop
	i := 0
	for {
		fmt.Println(i)
		i++
		if i > 5 {
			break
		}
	}

}
