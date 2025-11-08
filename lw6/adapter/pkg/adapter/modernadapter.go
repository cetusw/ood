package adapter

import (
	"adapter/pkg/graphicslib"
	"adapter/pkg/model"
	"adapter/pkg/moderngraphicslib"
)

type modernRendererAdapter struct {
	*moderngraphicslib.ModernGraphicsRenderer
	currentPos model.Point
	color      uint32
}

func NewModernRendererAdapter(renderer *moderngraphicslib.ModernGraphicsRenderer) graphicslib.Canvas {
	return &modernRendererAdapter{
		ModernGraphicsRenderer: renderer,
		currentPos:             model.Point{X: 0, Y: 0},
	}
}

func (a *modernRendererAdapter) SetColor(color uint32) {
	a.color = color
}

func (a *modernRendererAdapter) MoveTo(x, y int) {
	a.currentPos = model.Point{X: x, Y: y}
}

func (a *modernRendererAdapter) LineTo(x, y int) {
	endPoint := model.Point{X: x, Y: y}
	rgba := model.Uint32ToColor(a.color)
	a.DrawLine(a.currentPos, endPoint, rgba)
	a.currentPos = endPoint
}
