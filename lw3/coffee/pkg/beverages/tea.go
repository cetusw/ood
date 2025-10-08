package beverages

type Tea struct {
	beverage
}

func NewTea() *Tea {
	return &Tea{beverage{description: "Tea"}}
}

func (t *Tea) GetCost() float64 {
	return 30
}
