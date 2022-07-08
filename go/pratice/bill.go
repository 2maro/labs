package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

//new bill
func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{
			"cake": 5.99,
			"pie":  3.99,
		},
		tip: 0,
	}
	return b
}

//format bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf(" %-15v ....$%v \n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("%-15v ....$%v\n", " tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-15v ....$%0.2f", " total:", total+b.tip)

	return fs
}

//update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

//add items
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func main() {
	foobill := newBill("foo's bill")

	foobill.addItem("veg pie", 8.56)
	foobill.addItem("apple pie", 3.6)
	foobill.addItem("coffee", 1.56)

	foobill.updateTip(1.25)

	fmt.Println(foobill.format())

}
