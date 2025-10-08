package beverages

type Beverage interface {
	GetDescription() string
	GetCost() float64
}

type beverage struct {
	description string
}

func (b *beverage) GetDescription() string {
	return b.description
}
