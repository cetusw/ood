package shapedrawinglib

import (
	"adapter/pkg/graphicslib"
	"adapter/pkg/model"
)

type Rectangle struct {
	leftTop model.Point
	width   int
	height  int
}

func NewRectangle(leftTop model.Point, width, height int) *Rectangle {
	return &Rectangle{leftTop, width, height}
}

func (r *Rectangle) Draw(canvas graphicslib.Canvas) {
	p2 := model.Point{X: r.leftTop.X + r.width, Y: r.leftTop.Y}
	p3 := model.Point{X: r.leftTop.X + r.width, Y: r.leftTop.Y + r.height}
	p4 := model.Point{X: r.leftTop.X, Y: r.leftTop.Y + r.height}

	canvas.MoveTo(r.leftTop.X, r.leftTop.Y)
	canvas.LineTo(p2.X, p2.Y)
	canvas.LineTo(p3.X, p3.Y)
	canvas.LineTo(p4.X, p4.Y)
	canvas.LineTo(r.leftTop.X, r.leftTop.Y)
}
