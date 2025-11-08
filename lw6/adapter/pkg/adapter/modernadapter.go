package adapter

import (
	"adapter/pkg/graphicslib"
	"adapter/pkg/model"
	"adapter/pkg/moderngraphicslib"
)

type modernRendererAdapter struct {
	renderer   *moderngraphicslib.ModernGraphicsRenderer
	currentPos model.Point
}

func NewModernRendererAdapter(renderer *moderngraphicslib.ModernGraphicsRenderer) graphicslib.Canvas {
	return &modernRendererAdapter{
		renderer:   renderer,
		currentPos: model.Point{X: 0, Y: 0},
	}
}

func (a *modernRendererAdapter) MoveTo(x, y int) {
	a.currentPos = model.Point{X: x, Y: y}
}

func (a *modernRendererAdapter) LineTo(x, y int) {
	endPoint := model.Point{X: x, Y: y}
	a.renderer.DrawLine(a.currentPos, endPoint)
	a.currentPos = endPoint
}
