package main

import "fmt"

//Struct work ass classes in Go
//they create a blue print for data

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bill func
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// Format the bill func
func (b *bill) format() string {

	fs := "Bill breakdown : \n"
	var total float64 = 0

	//list the bill items

	for k, v := range b.items {
		fs += fmt.Sprintf("%-20v ....$%v \n", k+":", v)
		total += v
	}

	//add Tip

	fs += fmt.Sprintf("%-25v ...$%v \n", "tip:", b.tip)

	//add Total
	fs += fmt.Sprintf("%-20v ... $%0.2f", "total", total+b.tip)

	return fs

}

// update a tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

//add items to the bill

func (b bill) addItem(name string, price float64) {
	b.items[name] = price
}

//save the bill
