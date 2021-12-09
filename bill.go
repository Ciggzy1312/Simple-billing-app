package main

import "fmt"

// custom type using struct
type bill struct {
	name  string
	items map[string]float64
	tip   int
}

func newBill(name string) bill {
	bill := bill{
		name:  name,
		items: map[string]float64{"pie": 4.66, "coffee": 2.5, "soup": 3.2},
		tip:   0,
	}

	return bill
}

func (b bill) format() string {
	fs := "Bill breakdown \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-15v  $%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-15v  $%0.2f", "total:", total)

	return fs
}
