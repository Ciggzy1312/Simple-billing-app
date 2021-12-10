package main

import (
	"fmt"
	"os"
)

// custom type using struct
type bill struct {
	name  string
	items map[string]float64
	tip   int
}

func newBill(name string) bill {
	bill := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return bill
}

func (b *bill) updateTip(tip int) {
	(*b).tip = tip
}

func (b *bill) addItems(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill saved")
}

func (b bill) format() string {
	fs := "Bill breakdown \n"
	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-15v  $%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-15v  $%d\n", "tip:", b.tip)

	fs += fmt.Sprintf("%-15v  $%0.2f", "total:", total+float64(b.tip))

	return fs
}
