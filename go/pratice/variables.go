package main

import "fmt"

func main() {

	var a int
	var b string
	var c float64
	var d bool

	fmt.Println("this is a var's zero value", a, b, c, d)
	fmt.Printf("var a \t %T [%v]\n", a, a)
	fmt.Printf("var b \t %T [%v]\n", b, b)
	fmt.Printf("var c \t %T [%v]\n", c, c)
	fmt.Printf("var d \t %T [%v]\n", d, d)

}
