package shapes

import (
	"factory/pkg/canvas"
	"factory/pkg/model"
)

type rectangle struct {
	baseShape
	leftTop     model.Point
	rightBottom model.Point
}

func NewRectangle(color model.Color, p1 model.Point, p2 model.Point) Shape {
	return &rectangle{baseShape: baseShape{color: color}, leftTop: p1, rightBottom: p2}
}

func (r *rectangle) Draw(canvas canvas.Canvas) {
	canvas.SetColor(r.color)
	p1 := r.leftTop
	p2 := model.Point{X: r.rightBottom.X, Y: r.leftTop.Y}
	p3 := r.rightBottom
	p4 := model.Point{X: r.leftTop.X, Y: r.rightBottom.Y}
	canvas.DrawLine(p1, p2)
	canvas.DrawLine(p2, p3)
	canvas.DrawLine(p3, p4)
	canvas.DrawLine(p4, p1)
}
