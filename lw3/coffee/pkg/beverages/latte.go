package beverages

import "coffee/pkg/model"

type Latte struct {
	Coffee
	portion model.PortionType
}

func NewLatte(portion model.PortionType) *Latte {
	l := &Latte{}
	l.portion = portion
	l.description = "Latte" + " " + string(portion)
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
