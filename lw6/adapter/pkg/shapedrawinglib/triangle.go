package shapedrawinglib

import (
	"adapter/pkg/graphicslib"
	"adapter/pkg/model"
)

type Triangle struct {
	p1, p2, p3 model.Point
}

func NewTriangle(p1, p2, p3 model.Point) *Triangle {
	return &Triangle{p1, p2, p3}
}

func (t *Triangle) Draw(canvas graphicslib.Canvas) {
	canvas.MoveTo(t.p1.X, t.p1.Y)
	canvas.LineTo(t.p2.X, t.p2.Y)
	canvas.LineTo(t.p3.X, t.p3.Y)
	canvas.LineTo(t.p1.X, t.p1.Y)
}
