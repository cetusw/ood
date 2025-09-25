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

func (s *TriangleStrategy) Draw(canvas Canvas, id string, color string) string {
	return fmt.Sprintf(
		"%s drawn: id: <%s>, color <%s>, vertex1 <%.2f, %.2f>, vertex2 <%.2f, %.2f>, vertex3 <%.2f, %.2f>",
		common.Triangle,
		id,
		color,
		s.Vertices[0].X,
		s.Vertices[0].Y,
		s.Vertices[1].X,
		s.Vertices[1].Y,
		s.Vertices[2].X,
		s.Vertices[2].Y,
	)
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
