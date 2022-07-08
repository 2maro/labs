package main

import "fmt"

func updateName(x *string) {
	*x = "bob"
}
func main() {
	name := "Tom"

	fmt.Println(name)

	t := &name

	updateName(t)
	fmt.Println(name)
}
