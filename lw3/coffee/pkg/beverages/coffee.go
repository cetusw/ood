package beverages

type Coffee struct {
	beverage
}

func NewCoffee(description string) *Coffee {
	if description == "" {
		description = "Coffee"
	}
	return &Coffee{beverage{description: description}}
}

func (c *Coffee) GetCost() float64 {
	return 60
}
