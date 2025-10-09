package condiments

import (
	"coffee/pkg/beverages"
	"coffee/pkg/model"
)

const liquorCost = 50

type Liquor struct {
	CondimentDecorator
	liquorType model.LiquorType
}

func NewLiquor(beverage beverages.Beverage, liquorType model.LiquorType) *Liquor {
	return &Liquor{
		CondimentDecorator: CondimentDecorator{beverage: beverage},
		liquorType:         liquorType,
	}
}

func (l *Liquor) GetDescription() string {
	return l.beverage.GetDescription() + ", Liquor " + string(l.liquorType)
}

func (l *Liquor) GetCost() float64 {
	return l.beverage.GetCost() + liquorCost
}
