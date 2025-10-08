package condiments

import (
	"coffee/pkg/beverages"
	"strconv"
)

type Lemon struct {
	CondimentDecorator
	quantity uint
}

func NewLemon(beverage beverages.Beverage, quantity uint) *Lemon {
	if quantity == 0 {
		quantity = 1
	}
	return &Lemon{
		CondimentDecorator: CondimentDecorator{beverage: beverage},
		quantity:           quantity,
	}
}

func (l *Lemon) GetDescription() string {
	return l.beverage.GetDescription() + ", Lemon x " + strconv.FormatUint(uint64(l.quantity), 10)
}

func (l *Lemon) GetCost() float64 {
	return l.beverage.GetCost() + (10.0 * float64(l.quantity))
}
