package graphicslib

import "fmt"

type Canvas interface {
	MoveTo(x, y int)
	LineTo(x, y int)
}

type canvas struct{}

func NewCanvas() Canvas {
	return &canvas{}
}

func (c *canvas) MoveTo(x, y int) {
	fmt.Printf("MoveTo (%d, %d)\n", x, y)
}

func (c *canvas) LineTo(x, y int) {
	fmt.Printf("LineTo (%d, %d)\n", x, y)
}
