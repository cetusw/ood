package condiments

import (
	"coffee/pkg/beverages"
)

const creamCost = 25

type Cream struct {
	CondimentDecorator
}

func NewCream(beverage beverages.Beverage) *Cream {
	return &Cream{
		CondimentDecorator: CondimentDecorator{beverage: beverage},
	}
}

func (l *Cream) GetDescription() string {
	return l.beverage.GetDescription() + ", Cream"
}

func (l *Cream) GetCost() float64 {
	return l.beverage.GetCost() + creamCost
}
