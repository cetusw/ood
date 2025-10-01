package shape

import (
	"fmt"
	"shapes/pkg/common"
	"shapes/pkg/shapes/model"
)

type RectangleStrategy struct {
	topLeftPoint model.Point
	width        float64
	height       float64
}

func NewRectangleStrategy(
	topLeftPoint model.Point,
	width float64,
	height float64,
) *RectangleStrategy {
	return &RectangleStrategy{
		topLeftPoint: topLeftPoint,
		width:        width,
		height:       height,
	}
}

func (s *RectangleStrategy) Draw(canvas Canvas, id string, color string) string {
	return fmt.Sprintf(
		"rectangle drawn: id: <%s>, color <%s>, topLeftPoint <%.2f, %.2f>, width <%.2f>, height <%.2f>",
		id,
		color,
		s.topLeftPoint.X,
		s.topLeftPoint.Y,
		s.width,
		s.height,
	)
}

func (s *RectangleStrategy) MoveShape(vector model.Point) {
	s.topLeftPoint.X = s.topLeftPoint.X + vector.X
	s.topLeftPoint.Y = s.topLeftPoint.Y + vector.Y
}

func (s *RectangleStrategy) GetShapeInfo() string {
	return fmt.Sprintf(
		"%s %.2f %.2f %.2f %.2f",
		common.Rectangle,
		s.topLeftPoint.X,
		s.topLeftPoint.Y,
		s.width,
		s.height,
	)
}
