package shapes

import (
	"fmt"
	"shapes/pkg/shapes/model"
	"shapes/pkg/shapes/shape"
)

type Picture struct {
	shapes []*shape.Shape
	canvas shape.Canvas
}

func NewPicture() *Picture {
	return &Picture{
		shapes: []*shape.Shape{},
		canvas: shape.NewCanvas(800, 800),
	}
}

func (p *Picture) AddShape(shape *shape.Shape) error {
	for _, s := range p.shapes {
		if s.GetID() == shape.GetID() {
			return fmt.Errorf("фигура с таким названием уже существует")
		}
	}
	p.shapes = append(p.shapes, shape)
	return nil
}

func (p *Picture) MoveShape(shapeID string, vector model.Point) {
	for _, s := range p.shapes {
		if s.GetID() == shapeID {
			s.GetStrategy().MoveShape(vector)
		}
	}
}

func (p *Picture) MovePicture(vector model.Point) {
	for _, s := range p.shapes {
		s.GetStrategy().MoveShape(vector)
	}
}

func (p *Picture) DeleteShape(shapeID string) {
	for i := len(p.shapes) - 1; i >= 0; i-- {
		if p.shapes[i].GetID() == shapeID {
			p.shapes = append(p.shapes[:i], p.shapes[i+1:]...)
		}
	}
}

func (p *Picture) ListShapes() {
	for i, s := range p.shapes {
		fmt.Println(fmt.Sprintf(
			"%d %s %s %s",
			i,
			s.GetID(),
			s.GetColor(),
			s.GetStrategy().GetShapeInfo(),
		))
	}
}

func (p *Picture) ChangeColor(shapeID string, color string) {
	for _, s := range p.shapes {
		if s.GetID() == shapeID {
			s.SetColor(color)
		}
	}
}

func (p *Picture) ChangeShape(shapeID string, newStrategy shape.Strategy) {
	for _, s := range p.shapes {
		if s.GetID() == shapeID {
			s.SetStrategy(newStrategy)
		}
	}
}

func (p *Picture) DrawShape(shapeID string, canvas shape.Canvas) {
	for _, s := range p.shapes {
		if s.GetID() == shapeID {
			s.GetStrategy().Draw(canvas, s.GetColor())
		}
	}
}

func (p *Picture) DrawPicture(canvas shape.Canvas) {
	for _, s := range p.shapes {
		s.GetStrategy().Draw(canvas, s.GetColor())
	}
}

func (p *Picture) GetCanvas() shape.Canvas {
	return p.canvas
}

func (p *Picture) GetShapes() []*shape.Shape {
	return p.shapes
}
