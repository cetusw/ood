package beverages

type Latte struct {
	Coffee
}

func NewLatte() *Latte {
	l := &Latte{}
	l.description = "Latte"
	return l
}

func (l *Latte) GetCost() float64 {
	return 90
}
