package main

import "fmt"

type example struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	var e1 example

	fmt.Printf("%+v\n", e1)

	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.145,
	}

	fmt.Printf("\n%+v\n", e2)
	fmt.Printf("\n%+v\n", e2.counter)
	fmt.Printf("\n%+v\n", e2.pi)

	//on  the fly declaration
	e3 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.145,
	}
	fmt.Printf("\n%+v\n", e3)
}
