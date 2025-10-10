package shapes

import (
	"factory/pkg/canvas"
	"factory/pkg/model"
)

type triangle struct {
	baseShape
	v1, v2, v3 model.Point
}

func NewTriangle(color model.Color, v1, v2, v3 model.Point) Shape {
	return &triangle{baseShape: baseShape{color: color}, v1: v1, v2: v2, v3: v3}
}

func (t *triangle) Draw(canvas canvas.Canvas) {
	canvas.SetColor(t.color)
	canvas.DrawLine(t.v1, t.v2)
	canvas.DrawLine(t.v2, t.v3)
	canvas.DrawLine(t.v3, t.v1)
}
