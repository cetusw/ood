package shape

import (
	"fmt"
	"shapes/pkg/common"
	"shapes/pkg/shapes/model"
)

type LineStrategy struct {
	Vertices [2]model.Point
}

func NewLineStrategy(vertices [2]model.Point) *LineStrategy {
	return &LineStrategy{
		Vertices: vertices,
	}
}

func (s *LineStrategy) Draw(canvas Canvas, color string) {
	canvas.SetColor(color)
	canvas.MoveTo(s.Vertices[0])
	canvas.LineTo(s.Vertices[1])
}

func (s *LineStrategy) MoveShape(vector model.Point) {
	for vertex := range s.Vertices {
		s.Vertices[vertex].X = s.Vertices[vertex].X + vector.X
		s.Vertices[vertex].Y = s.Vertices[vertex].Y + vector.Y
	}
}

func (s *LineStrategy) GetShapeInfo() string {
	return fmt.Sprintf(
		"%s %.2f %.2f %.2f %.2f",
		common.Line,
		s.Vertices[0].X,
		s.Vertices[0].Y,
		s.Vertices[1].X,
		s.Vertices[1].Y,
	)
}
