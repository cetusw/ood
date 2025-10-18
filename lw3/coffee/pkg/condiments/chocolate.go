package condiments

import (
	"coffee/pkg/beverages"
	"fmt"
)

const chocolateCost = 10

type Chocolate struct {
	CondimentDecorator
	quantity int // TODO: uint
}

func NewChocolate(beverage beverages.Beverage, quantity int) *Chocolate {
	if quantity > 5 {
		quantity = 5
	}
	return &Chocolate{
		CondimentDecorator: CondimentDecorator{beverage: beverage},
		quantity:           quantity,
	}
}

func (l *Chocolate) GetDescription() string {
	return fmt.Sprintf("%s %d", l.beverage.GetDescription()+", Chocolate x ", l.quantity)
}

func (l *Chocolate) GetCost() float64 {
	return l.beverage.GetCost() + float64(chocolateCost*l.quantity)
}
