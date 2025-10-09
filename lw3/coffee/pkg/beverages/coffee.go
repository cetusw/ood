package beverages

type Coffee struct {
	beverage
}

func NewCoffee() *Coffee {
	return &Coffee{
		beverage: beverage{
			description: "Coffee",
		},
	}
}

func (c *Coffee) GetCost() float64 {
	return 60
}
