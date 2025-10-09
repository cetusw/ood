package main

import (
	"bufio"
	"coffee/pkg/beverages"
	"coffee/pkg/condiments"
	"coffee/pkg/model"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readChoice(reader *bufio.Reader) int {
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return choice
}

func selectPortion(reader *bufio.Reader) model.PortionType {
	fmt.Println("Choose portion: 1 - Regular, 2 - Double")
	choice := readChoice(reader)
	if choice == 2 {
		return model.Double
	}
	return model.Regular
}

func selectSize(reader *bufio.Reader) model.SizeType {
	fmt.Println("Choose size: 1 - Small, 2 - Middle, 3 - Large")
	choice := readChoice(reader)
	switch choice {
	case 1:
		return model.Small
	case 2:
		return model.Middle
	case 3:
		return model.Large
	default:
		return model.Small
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	portionDescriber := model.NewPortionDescriber()
	sizeDescriber := model.NewSizeDescriber()

	var beverage beverages.Beverage

	fmt.Println("Type to choose beverage: 1 - Coffee, 2 - Milkshake, 3 - Tea")
	mainChoice := readChoice(reader)

	switch mainChoice {
	case 1:
		fmt.Println("Specify which one: 1 - Cappuccino, 2 - Latte, 3 - Plain Coffee")
		coffeeChoice := readChoice(reader)

		switch coffeeChoice {
		case 1:
			selectedPortion := selectPortion(reader)
			beverage = beverages.NewCappuccino(selectedPortion, portionDescriber)
		case 2:
			selectedPortion := selectPortion(reader)
			beverage = beverages.NewLatte(selectedPortion, portionDescriber)
		case 3:
			beverage = beverages.NewCoffee()
		default:
			return
		}
	case 2:
		selectedSize := selectSize(reader)
		beverage = beverages.NewMilkshake(selectedSize, sizeDescriber)
	case 3:
		beverage = beverages.NewTea()

	default:
		return
	}

	if beverage == nil {
		return
	}

	for {
		fmt.Println("1 - Lemon, 2 - Cinnamon, 0 - Checkout")
		condimentChoice := readChoice(reader)

		switch condimentChoice {
		case 1:
			beverage = condiments.NewLemon(beverage, 2)
		case 2:
			beverage = condiments.NewCinnamon(beverage)
		case 0:
			fmt.Printf("%s, cost: %.2f\n", beverage.GetDescription(), beverage.GetCost())
			return
		default:
			continue
		}
	}
}
