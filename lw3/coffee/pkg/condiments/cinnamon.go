package condiments

import "coffee/pkg/beverages"

type Cinnamon struct {
	CondimentDecorator
}

func NewCinnamon(beverage beverages.Beverage) *Cinnamon {
	return &Cinnamon{CondimentDecorator{beverage: beverage}}
}

func (c *Cinnamon) GetDescription() string {
	return c.beverage.GetDescription() + ", Cinnamon"
}

func (c *Cinnamon) GetCost() float64 {
	return c.beverage.GetCost() + 20
}
