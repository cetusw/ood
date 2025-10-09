package canvas

import (
	"fmt"

	"factory/pkg/domain"
)

type canvas struct{}

func NewCanvas() domain.Canvas {
	return &canvas{}
}

func (c *canvas) SetColor(color domain.Color) {
	fmt.Printf("SetColor: %s\n", color)
}

func (c *canvas) DrawLine(from, to domain.Point) {
	fmt.Printf("DrawLine: from (%d, %d) to (%d, %d)\n", from.X, from.Y, to.X, to.Y)
}

func (c *canvas) DrawEllipse(center domain.Point, hRadius, vRadius int) {
	fmt.Printf("DrawEllipse: center (%d, %d), hRadius %d, vRadius %d\n", center.X, center.Y, hRadius, vRadius)
}
