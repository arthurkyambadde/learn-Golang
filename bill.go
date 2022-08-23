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
		name: name,
		items: map[string]float64{
			"katogo": 2.99,
			"rice":   10.65,
			"irish":  5.22,
		},
		tip: 0,
	}
	return b
}

// Format the bill func
func (b bill) format() string {

	fs := "Bill breakdown : \n"
	var total float64 = 0

	//list the bill items

	for k, v := range b.items {
		fs += fmt.Sprintf("%-20v ....$%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-20v ... $%0.2f", "total", total)

	return fs

}
