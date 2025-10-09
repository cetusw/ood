package beverages

import "coffee/pkg/model"

type Latte struct {
	Coffee
	portion model.PortionType
}

func NewLatte(portion model.PortionType, describer model.PortionDescriber) *Latte {
	l := &Latte{}
	l.portion = portion
	l.description = "Latte" + describer.GetDescriptionPrefix(portion)
	return l
}

func (l *Latte) GetCost() float64 {
	switch l.portion {
	case model.Double:
		return 130
	default:
		return 90
	}
}
