package shapes

import "factory/pkg/domain"

type ellipse struct {
	baseShape
	center           domain.Point
	hRadius, vRadius int
}

func NewEllipse(color domain.Color, center domain.Point, hRadius, vRadius int) domain.Shape {
	return &ellipse{baseShape: baseShape{color: color}, center: center, hRadius: hRadius, vRadius: vRadius}
}

func (e *ellipse) Draw(canvas domain.Canvas) {
	canvas.SetColor(e.color)
	canvas.DrawEllipse(e.center, e.hRadius, e.vRadius)
}
