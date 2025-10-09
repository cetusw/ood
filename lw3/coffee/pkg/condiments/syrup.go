package condiments

import (
	"coffee/pkg/beverages"
	"coffee/pkg/model"
)

const syrupCost = 15

type Syrup struct {
	CondimentDecorator
	syrupType model.SyrupType
}

func NewSyrup(beverage beverages.Beverage, syrupType model.SyrupType) *Syrup {
	return &Syrup{
		CondimentDecorator: CondimentDecorator{beverage: beverage},
		syrupType:          syrupType,
	}
}

func (s *Syrup) GetDescription() string {
	return s.beverage.GetDescription() + ", " + string(s.syrupType) + " syrup"
}

func (s *Syrup) GetCost() float64 {
	return s.beverage.GetCost() + syrupCost
}
