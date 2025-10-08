package condiments

import "coffee/pkg/beverages"

type SyrupType int

const (
	ChocolateSyrup SyrupType = iota
	MapleSyrup
)

type Syrup struct {
	CondimentDecorator
	syrupType SyrupType
}

func NewSyrup(beverage beverages.Beverage, syrupType SyrupType) *Syrup {
	return &Syrup{
		CondimentDecorator: CondimentDecorator{beverage: beverage},
		syrupType:          syrupType,
	}
}

func (s *Syrup) GetDescription() string {
	typeStr := "Chocolate"
	if s.syrupType == MapleSyrup {
		typeStr = "Maple"
	}
	return s.beverage.GetDescription() + ", " + typeStr + " syrup"
}

func (s *Syrup) GetCost() float64 {
	return s.beverage.GetCost() + 15
}
