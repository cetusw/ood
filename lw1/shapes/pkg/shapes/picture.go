package shapes

import (
	"fmt"
	"shapes/pkg/shapes/shape"
)

type Picture struct {
	shapes []*shape.Shape
}

func NewPicture() *Picture {
	return &Picture{shapes: []*shape.Shape{}}
}

func (p *Picture) AddShape(shape *shape.Shape) error {
	for _, s := range p.shapes {
		if s.Id == shape.Id {
			return fmt.Errorf("фигура с таким названием уже существует")
		}
	}
	p.shapes = append(p.shapes, shape)
	return nil
}

func (p *Picture) ListShapes() {
	for i, s := range p.shapes {
		fmt.Println(fmt.Sprintf("%d %s %s %s", i, s.Id, s.Color, s.Strategy.GetShapeInfo()))
	}
}
