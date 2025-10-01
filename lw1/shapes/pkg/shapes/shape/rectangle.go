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

func (s *RectangleStrategy) Draw(canvas Canvas, color string) {
	canvas.SetColor(color)
	canvas.MoveTo(s.topLeftPoint)
	canvas.LineTo(common.MovePoint(s.topLeftPoint, model.Point{
		X: s.width,
		Y: 0,
	}))
	canvas.LineTo(common.MovePoint(s.topLeftPoint, model.Point{
		X: s.width,
		Y: -s.height,
	}))
	canvas.LineTo(common.MovePoint(s.topLeftPoint, model.Point{
		X: 0,
		Y: -s.height,
	}))
	canvas.LineTo(s.topLeftPoint)
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
