package beverages

type Cappuccino struct {
	Coffee
}

func NewCappuccino() *Cappuccino {
	c := &Cappuccino{}
	c.description = "Cappuccino"
	return c
}

func (c *Cappuccino) GetCost() float64 {
	return 80
}
