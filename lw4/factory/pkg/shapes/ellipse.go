package shapes

import (
	"factory/pkg/canvas"
	"factory/pkg/model"
)

type ellipse struct {
	baseShape
	center           model.Point
	hRadius, vRadius int
}

func NewEllipse(color model.Color, center model.Point, hRadius, vRadius int) Shape {
	return &ellipse{baseShape: baseShape{color: color}, center: center, hRadius: hRadius, vRadius: vRadius}
}

func (e *ellipse) Draw(canvas canvas.Canvas) {
	canvas.SetColor(e.color)
	canvas.DrawEllipse(e.center, e.hRadius, e.vRadius)
}
