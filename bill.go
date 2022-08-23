package main

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
