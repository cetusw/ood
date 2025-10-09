package shapes

import "factory/pkg/domain"

type rectangle struct {
	baseShape
	leftTop     domain.Point
	rightBottom domain.Point
}

func NewRectangle(color domain.Color, p1 domain.Point, p2 domain.Point) domain.Shape {
	return &rectangle{baseShape: baseShape{color: color}, leftTop: p1, rightBottom: p2}
}

func (r *rectangle) Draw(canvas domain.Canvas) {
	canvas.SetColor(r.color)
	p1 := r.leftTop
	p2 := domain.Point{X: r.rightBottom.X, Y: r.leftTop.Y}
	p3 := r.rightBottom
	p4 := domain.Point{X: r.leftTop.X, Y: r.rightBottom.Y}
	canvas.DrawLine(p1, p2)
	canvas.DrawLine(p2, p3)
	canvas.DrawLine(p3, p4)
	canvas.DrawLine(p4, p1)
}
