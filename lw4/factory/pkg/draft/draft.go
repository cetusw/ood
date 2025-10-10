package draft

import (
	"factory/pkg/canvas"
	"factory/pkg/shapes"
)

type Draft struct {
	shapes []shapes.Shape
}

func (p *Draft) AddShape(s shapes.Shape) {
	p.shapes = append(p.shapes, s)
}

func (p *Draft) Draw(canvas canvas.Canvas) {
	for _, shape := range p.shapes {
		shape.Draw(canvas)
	}
}

func (p *Draft) GetShapeCount() int {
	return len(p.shapes)
}
