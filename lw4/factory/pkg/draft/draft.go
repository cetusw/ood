package draft

import (
	"factory/pkg/canvas"
	"factory/pkg/shapes"
)

type Draft struct {
	shapes []shapes.Shape
}

func (d *Draft) AddShape(s shapes.Shape) {
	d.shapes = append(d.shapes, s)
}

func (d *Draft) Draw(canvas canvas.Canvas) {
	for _, shape := range d.shapes {
		shape.Draw(canvas)
	}
}

func (d *Draft) GetShapeCount() int {
	return len(d.shapes)
}
