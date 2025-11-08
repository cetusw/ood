package graphicslib

import (
	"fmt"
)

type Canvas interface {
	SetColor(color uint32)
	MoveTo(x, y int)
	LineTo(x, y int)
}

type canvas struct{}

func NewCanvas() Canvas {
	return &canvas{}
}

func (c *canvas) SetColor(color uint32) {
	fmt.Printf("SetColor (%d)\n", color)
}

func (c *canvas) MoveTo(x, y int) {
	fmt.Printf("MoveTo (%d, %d)\n", x, y)
}

func (c *canvas) LineTo(x, y int) {
	fmt.Printf("LineTo (%d, %d)\n", x, y)
}
