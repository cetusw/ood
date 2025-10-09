package model

type PortionDescriber interface {
	GetDescriptionPrefix(portion PortionType) string
}

type portionDescriber struct {
	portion map[PortionType]string
}

func NewPortionDescriber() PortionDescriber {
	return &portionDescriber{
		portion: map[PortionType]string{
			Double: "Double",
			Triple: "Triple",
		},
	}
}

func (pd *portionDescriber) GetDescriptionPrefix(portion PortionType) string {
	portionSize, ok := pd.portion[portion]
	if !ok {
		return ""
	}

	return " " + portionSize
}
