package condiments

import (
	"coffee/pkg/beverages"
	"strconv"
)

type CoconutFlakes struct {
	CondimentDecorator
	mass uint
}

func NewCoconutFlakes(beverage beverages.Beverage, mass uint) *CoconutFlakes {
	return &CoconutFlakes{
		CondimentDecorator: CondimentDecorator{beverage: beverage},
		mass:               mass,
	}
}

func (c *CoconutFlakes) GetDescription() string {
	return c.beverage.GetDescription() + ", Coconut flakes " + strconv.FormatUint(uint64(c.mass), 10) + "g"
}

func (c *CoconutFlakes) GetCost() float64 {
	return c.beverage.GetCost() + (1.0 * float64(c.mass))
}
