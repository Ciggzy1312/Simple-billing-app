package main

import (
	"bufio"
	"fmt"
	"os"
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

		fmt.Println(item, price)

		break

	case "t":
		tip, _ := getInput("Enter tip : ", reader)

		fmt.Println(tip)

	case "s":
		fmt.Println("You chose s")
		break

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

	mybill.addItems("crab meat soup", 3.2)
	mybill.addItems("shepherd's pie", 4.66)
	mybill.addItems("tiramisu", 2.5)

	mybill.updateTip(5)

	fmt.Println(mybill.format())
}
