package model

type SizeDescriber interface {
	GetSize(size SizeType) string
}

type sizeDescriber struct {
	descriptions map[SizeType]string
}

func NewSizeDescriber() SizeDescriber {
	return &sizeDescriber{
		descriptions: map[SizeType]string{
			Small:  "Small",
			Middle: "Middle",
			Large:  "Large",
		},
	}
}

func (sd *sizeDescriber) GetSize(size SizeType) string {
	stringSize, ok := sd.descriptions[size]
	if !ok {
		return ""
	}

	return " " + stringSize
}
