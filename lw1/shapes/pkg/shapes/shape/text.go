package shape

import (
	"fmt"
	"shapes/pkg/common"
	"shapes/pkg/shapes/model"
)

type TextStrategy struct {
	topLeftPoint model.Point
	fontSize     float64
	text         string
}

func NewTextStrategy(
	topLeftPoint model.Point,
	size float64,
	content string,
) *TextStrategy {
	return &TextStrategy{
		topLeftPoint: topLeftPoint,
		fontSize:     size,
		text:         content,
	}
}

func (s *TextStrategy) Draw(canvas Canvas, color string) {
	canvas.SetColor(color)
	canvas.MoveTo(s.topLeftPoint)
	canvas.DrawText(s.topLeftPoint, s.fontSize, s.text)
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
		s.fontSize,
		s.text,
	)
}
