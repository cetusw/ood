package beverages

import "coffee/pkg/model"

type Milkshake struct {
	beverage
	size      model.SizeType
	describer model.SizeDescriber
}

func NewMilkshake(size model.SizeType, describer model.SizeDescriber) *Milkshake {
	return &Milkshake{
		beverage: beverage{
			description: "Milkshake" + describer.GetSize(size),
		},
		size:      size,
		describer: describer,
	}
}

func (m *Milkshake) GetCost() float64 {
	switch m.size {
	case model.Small:
		return 50
	case model.Middle:
		return 60
	case model.Large:
		return 80
	default:
		return 50
	}
}
