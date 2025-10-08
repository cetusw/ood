package beverages

type Milkshake struct {
	beverage
}

func NewMilkshake() *Milkshake {
	return &Milkshake{beverage{description: "Milkshake"}}
}

func (m *Milkshake) GetCost() float64 {
	return 80
}
