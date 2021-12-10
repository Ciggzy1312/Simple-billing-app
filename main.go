package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose an option ->\n a -add item\n t - tip bill\n s - save bill\n ", reader)

	switch option {
	case "a":
		item, _ := getInput("Enter name of item : ", reader)

		price, _ := getInput("Enter price : ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}
		b.addItems(item, p)

		fmt.Printf("Item added - %v   Price - %v\n", item, price)
		promptOptions(b)

		break

	case "t":
		tip, _ := getInput("Enter tip : ", reader)

		t, err := strconv.Atoi(tip)
		if err != nil {
			fmt.Println("Tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Printf("Tip added - %v\n", tip)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("bill has been saved as", b.name)

	default:
		fmt.Println("Choose a valid option...")
	}

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	/*fmt.Print("Create a new bill name : ")

	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)*/

	name, _ := getInput("Create a new bill name : ", reader)

	b := newBill(name)

	return b
}

func main() {
	mybill := createBill()

	promptOptions(mybill)
}
