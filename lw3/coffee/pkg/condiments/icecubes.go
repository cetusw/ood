package condiments

import (
	"coffee/pkg/beverages"
	"coffee/pkg/model"
	"fmt"
)

const costPerWaterCube = 5
const costPerDryCube = 10

type IceCubes struct {
	CondimentDecorator
	quantity int
	iceType  model.IceCubeType
}

func NewIceCubes(beverage beverages.Beverage, quantity int, iceType model.IceCubeType) *IceCubes {
	return &IceCubes{
		CondimentDecorator: CondimentDecorator{beverage: beverage},
		quantity:           quantity,
		iceType:            iceType,
	}
}

func (i *IceCubes) GetDescription() string {
	return fmt.Sprintf("%s %d", i.beverage.GetDescription()+", "+string(i.iceType)+" ice cubes x ", i.quantity)
}

func (i *IceCubes) GetCost() float64 {
	costPerCube := costPerWaterCube
	if i.iceType == model.DryIce {
		costPerCube = costPerDryCube
	}
	return i.beverage.GetCost() + float64(costPerCube*i.quantity)
}
