package condiments

import (
	"coffee/pkg/beverages"
	"strconv"
)

type ChocolateCrumbs struct {
	CondimentDecorator
	mass uint
}

func NewChocolateCrumbs(beverage beverages.Beverage, mass uint) *ChocolateCrumbs {
	return &ChocolateCrumbs{
		CondimentDecorator: CondimentDecorator{beverage: beverage},
		mass:               mass,
	}
}

func (c *ChocolateCrumbs) GetDescription() string {
	return c.beverage.GetDescription() + ", Chocolate crumbs " + strconv.FormatUint(uint64(c.mass), 10) + "g"
}

func (c *ChocolateCrumbs) GetCost() float64 {
	return c.beverage.GetCost() + (2.0 * float64(c.mass))
}
