package main

import "fmt"

func main() {

	collection := make(map[string]int)

	collection["hello"] = 1
	collection["world"] = 2

	fmt.Println(collection)

	collection["trash"] = 3

	delete(collection, "trash")

	fmt.Println(collection)

}
