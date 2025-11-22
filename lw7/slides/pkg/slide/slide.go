package slide

import (
	"slices"
	"slides/pkg/canvas"
	"slides/pkg/model"
	"slides/pkg/shapes"
)

type Slide interface {
	InsertShape(shape shapes.Shape)
	RemoveShapeAtIndex(idx int)
	Draw(canvas canvas.Canvas)
	GetWidth() int
	GetHeight() int
	GetShapesCount() int
	GetShapeAtIndex(idx int) shapes.Shape
	GetBackgroundColor() model.Color
	SetBackgroundColor(color model.Color)
}

type slide struct {
	width           int
	height          int
	backgroundColor model.Color
	shapes          []shapes.Shape
}

func NewSlide(width int, height int) Slide {
	return &slide{
		width:  width,
		height: height,
	}
}

func (s *slide) InsertShape(shape shapes.Shape) {
	s.shapes = append(s.shapes, shape)
}

func (s *slide) RemoveShapeAtIndex(idx int) {
	s.shapes = slices.Delete(s.shapes, idx, idx+1)
}

func (s *slide) Draw(canvas canvas.Canvas) {
	for _, shape := range s.shapes {
		shape.Draw(canvas)
	}
}

func (s *slide) GetWidth() int {
	return s.width
}

func (s *slide) GetHeight() int {
	return s.height
}

func (s *slide) GetShapesCount() int {
	return len(s.shapes)
}

func (s *slide) GetShapeAtIndex(idx int) shapes.Shape {
	return s.shapes[idx]
}

func (s *slide) GetBackgroundColor() model.Color {
	return s.backgroundColor
}

func (s *slide) SetBackgroundColor(color model.Color) {
	s.backgroundColor = color
}
