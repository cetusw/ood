package shape

import (
	"fmt"
	"shapes/pkg/common"
	"shapes/pkg/shapes/model"
)

type TextStrategy struct {
	topLeftPoint model.Point
	Size         float64
	Content      string
}

func NewTextStrategy(
	topLeftPoint model.Point,
	size float64,
	content string,
) *TextStrategy {
	return &TextStrategy{
		topLeftPoint: topLeftPoint,
		Size:         size,
		Content:      content,
	}
}

func (s *TextStrategy) Draw(canvas Canvas, id string, color string) string {
	return fmt.Sprintf(
		"%s drawn: id: <%s>, color <%s>, topLeftPoint <%.2f, %.2f>, size <%.2f>, text <%s>",
		common.Text,
		id,
		color,
		s.topLeftPoint.X,
		s.topLeftPoint.Y,
		s.Size,
		s.Content,
	)
}

func (s *TextStrategy) MoveShape(vector model.Point) {
	s.topLeftPoint.X = s.topLeftPoint.X + vector.X
	s.topLeftPoint.Y = s.topLeftPoint.Y + vector.Y
}

func (s *TextStrategy) GetShapeInfo() string {
	return fmt.Sprintf(
		"%s %.2f %.2f %.2f %s",
		common.Text,
		s.topLeftPoint.X,
		s.topLeftPoint.Y,
		s.Size,
		s.Content,
	)
}
