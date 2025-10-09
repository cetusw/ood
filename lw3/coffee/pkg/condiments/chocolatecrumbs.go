package condiments

import (
	"coffee/pkg/beverages"
	"fmt"
)

const chocolateCrumbsCost = 2

type ChocolateCrumbs struct {
	CondimentDecorator
	mass int
}

func NewChocolateCrumbs(beverage beverages.Beverage, mass int) *ChocolateCrumbs {
	return &ChocolateCrumbs{
		CondimentDecorator: CondimentDecorator{beverage: beverage},
		mass:               mass,
	}
}

func (c *ChocolateCrumbs) GetDescription() string {
	return fmt.Sprintf("%s %d%s", c.beverage.GetDescription()+", Chocolate crumbs ", c.mass, "g")
}

func (c *ChocolateCrumbs) GetCost() float64 {
	return c.beverage.GetCost() + float64(chocolateCrumbsCost*c.mass)
}
