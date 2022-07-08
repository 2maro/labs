package main

import "fmt"

type Speak interface {
	Speak()
}

type Dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

func (d *Dog) Speak() {
	fmt.Printf("woof! My name is %s, it is %t I am a mammal with a pack fator of %v\n",
		d.Name,
		d.IsMammal,
		d.PackFactor,
	)
}

type Cat struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

func (c *Cat) Speak() {
	fmt.Printf("meow! My name is %v, it is %t I am a mammal with no oack factor of %v \n",
		c.Name,
		c.IsMammal,
		c.PackFactor,
	)
}
func main() {
	kitty := Cat{}
	kitty.Name("kitty")
	kitty.IsMammal(true)
	kitty.PackFactor(1)
	kitty.Speak()
}
