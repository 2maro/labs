package main

import "fmt"

type fighter struct {
	name  string
	style string
}

// value receiver
func (f fighter) status() {
	fmt.Printf("%v is a %v\n", f.name, f.style)
}

//pointer receiver
func (f *fighter) changestyle(style string) {
	f.style = style
}
func main() {

	joe := fighter{"joe", "bjj"}
	joe.changestyle("kickboxer")
	joe.status()
}
