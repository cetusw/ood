package beverages

import "coffee/pkg/model"

type Cappuccino struct {
	Coffee
	portion model.PortionType
}

func NewCappuccino(portion model.PortionType) *Cappuccino {
	c := &Cappuccino{}
	c.portion = portion
	c.description = "Cappuccino" + " " + string(portion)
	return c
}

func (c *Cappuccino) GetCost() float64 {
	switch c.portion {
	case model.Double:
		return 120
	default:
		return 80
	}
}
