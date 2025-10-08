package main

import (
	"bufio"
	"coffee/pkg/beverages"
	"coffee/pkg/condiments"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type 1 for coffee or 2 for tea")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var beverage beverages.Beverage

	switch input {
	case "1":
		beverage = beverages.NewCoffee("")
	case "2":
		beverage = beverages.NewTea()
	default:
		fmt.Println("Invalid choice.")
		return
	}

	for {
		fmt.Println("1 - Lemon, 2 - Cinnamon, 0 - Checkout")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			beverage = condiments.NewLemon(beverage, 2)
		case "2":
			beverage = condiments.NewCinnamon(beverage)
		case "0":
			fmt.Printf("%s, cost: %.2f\n", beverage.GetDescription(), beverage.GetCost())
			return
		default:
			fmt.Println("Invalid choice.")
			return
		}
	}
}
