package condiments

import (
	"coffee/pkg/beverages"
	"fmt"
)

const coconutFlakes = 1

type CoconutFlakes struct {
	CondimentDecorator
	mass int
}

func NewCoconutFlakes(beverage beverages.Beverage, mass int) *CoconutFlakes {
	return &CoconutFlakes{
		CondimentDecorator: CondimentDecorator{beverage: beverage},
		mass:               mass,
	}
}

func (c *CoconutFlakes) GetDescription() string {
	return fmt.Sprintf("%s %d%s", c.beverage.GetDescription()+", Coconut flakes ", c.mass, "g")
}

func (c *CoconutFlakes) GetCost() float64 {
	return c.beverage.GetCost() + float64(coconutFlakes*c.mass)
}
