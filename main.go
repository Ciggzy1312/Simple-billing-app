package main

import "fmt"

func main() {
	mybill := newBill("Mario")

	fmt.Println(mybill.format())

	fmt.Println(mybill)
}
