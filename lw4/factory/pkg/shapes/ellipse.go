package shapes

import (
	"factory/pkg/canvas"
	"factory/pkg/model"
)

type ellipse struct {
	baseShape
	center model.Point
	radius model.Radius
}

func NewEllipse(color model.Color, center model.Point, radius model.Radius) Shape {
	return &ellipse{
		baseShape: baseShape{color: color},
		center:    center,
		radius:    radius,
	}
}

func (e *ellipse) Draw(canvas canvas.Canvas) {
	canvas.SetColor(e.color)
	canvas.DrawEllipse(e.center, e.radius)
}
