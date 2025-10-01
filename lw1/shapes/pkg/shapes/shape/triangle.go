package shape

import (
	"fmt"
	"shapes/pkg/common"
	"shapes/pkg/shapes/model"
)

type TriangleStrategy struct {
	Vertices [3]model.Point
}

func NewTriangleStrategy(vertices [3]model.Point) *TriangleStrategy {
	return &TriangleStrategy{Vertices: vertices}
}

func (s *TriangleStrategy) Draw(canvas Canvas, color string) {
	canvas.SetColor(color)
	canvas.MoveTo(s.Vertices[0])
	canvas.LineTo(s.Vertices[1])
	canvas.LineTo(s.Vertices[2])
	canvas.LineTo(s.Vertices[0])
}

func (s *TriangleStrategy) MoveShape(vector model.Point) {
	for vertex := range s.Vertices {
		s.Vertices[vertex].X = s.Vertices[vertex].X + vector.X
		s.Vertices[vertex].Y = s.Vertices[vertex].Y + vector.Y
	}
}

func (s *TriangleStrategy) GetShapeInfo() string {
	return fmt.Sprintf(
		"%s %.2f %.2f %.2f %.2f %.2f %.2f",
		common.Triangle,
		s.Vertices[0].X,
		s.Vertices[0].Y,
		s.Vertices[1].X,
		s.Vertices[1].Y,
		s.Vertices[2].X,
		s.Vertices[2].Y,
	)
}
