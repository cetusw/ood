package beverages

import "coffee/pkg/model"

type Tea struct {
	beverage
}

func NewTea(teaType model.TeaType) *Tea {
	return &Tea{
		beverage{
			description: "Tea" + " " + string(teaType),
		},
	}
}

func (t *Tea) GetCost() float64 {
	return 30
}
