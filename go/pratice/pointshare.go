package main

import "fmt"

func main() {
	x := 10

	//data is the value itself "x" and the address of that value
	fmt.Println("count:\tValue of [", x, "]\tAddr of [", &x, "]")

	incr(&x)

	fmt.Println("count:\tValue of [", x, "]\tAddr of [", &x, "]")
}

func incr(i *int) int {
	*i++
	fmt.Println("count:\tValue of [", i, "]\tAddr of [", &i, "]")
	return *i
}
