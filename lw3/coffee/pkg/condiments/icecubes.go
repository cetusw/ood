package condiments

import (
	"coffee/pkg/beverages"
	"strconv"
)

type IceCubeType int

const (
	DryIce IceCubeType = iota
	WaterIce
)

type IceCubes struct {
	CondimentDecorator
	quantity uint
	iceType  IceCubeType
}

func NewIceCubes(beverage beverages.Beverage, quantity uint, iceType IceCubeType) *IceCubes {
	return &IceCubes{
		CondimentDecorator: CondimentDecorator{beverage: beverage},
		quantity:           quantity,
		iceType:            iceType,
	}
}

func (i *IceCubes) GetDescription() string {
	typeStr := "Water"
	if i.iceType == DryIce {
		typeStr = "Dry"
	}
	return i.beverage.GetDescription() + ", " + typeStr + " ice cubes x " + strconv.FormatUint(uint64(i.quantity), 10)
}

func (i *IceCubes) GetCost() float64 {
	costPerCube := 5.0
	if i.iceType == DryIce {
		costPerCube = 10.0
	}
	return i.beverage.GetCost() + (costPerCube * float64(i.quantity))
}
