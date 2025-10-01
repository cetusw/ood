package shapes

import (
	"fmt"
	"shapes/pkg/shapes/model"
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
		if s.GetId() == shape.GetId() {
			return fmt.Errorf("фигура с таким названием уже существует")
		}
	}
	p.shapes = append(p.shapes, shape)
	return nil
}

func (p *Picture) MoveShape(id string, vector model.Point) {
	for _, s := range p.shapes {
		if s.GetId() == id {
			s.GetStrategy().MoveShape(vector)
		}
	}
}

func (p *Picture) MovePicture(vector model.Point) {
	for _, s := range p.shapes {
		s.GetStrategy().MoveShape(vector)
	}
}

func (p *Picture) DeleteShape(id string) {
	for i := len(p.shapes) - 1; i >= 0; i-- {
		if p.shapes[i].GetId() == id {
			p.shapes = append(p.shapes[:i], p.shapes[i+1:]...)
		}
	}
}

func (p *Picture) ListShapes() {
	for i, s := range p.shapes {
		fmt.Println(fmt.Sprintf(
			"%d %s %s %s",
			i,
			s.GetId(),
			s.GetColor(),
			s.GetStrategy().GetShapeInfo(),
		))
	}
}

func (p *Picture) ChangeColor(id string, color string) {
	for _, s := range p.shapes {
		if s.GetId() == id {
			s.SetColor(color)
		}
	}
}

func (p *Picture) ChangeShape(id string, newStrategy shape.Strategy) {
	for _, s := range p.shapes {
		if s.GetId() == id {
			s.SetStrategy(newStrategy)
		}
	}
}
