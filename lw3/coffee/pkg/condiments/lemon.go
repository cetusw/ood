package condiments

import (
	"coffee/pkg/beverages"
	"fmt"
)

const lemonCost = 10

type Lemon struct {
	CondimentDecorator
	quantity int
}

func NewLemon(beverage beverages.Beverage, quantity int) *Lemon {
	if quantity == 0 {
		quantity = 1
	}
	return &Lemon{
		CondimentDecorator: CondimentDecorator{beverage: beverage},
		quantity:           quantity,
	}
}

func (l *Lemon) GetDescription() string {
	return fmt.Sprintf("%s %d", l.beverage.GetDescription()+", Lemon x ", l.quantity)
}

func (l *Lemon) GetCost() float64 {
	return l.beverage.GetCost() + (lemonCost * float64(l.quantity))
}
