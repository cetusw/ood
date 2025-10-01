package shape

import (
	"fmt"
	"shapes/pkg/common"
	"shapes/pkg/shapes/model"
)

type CircleStrategy struct {
	Center model.Point
	Radius model.Radius
}

func NewCircleStrategy(center model.Point, radius float64) *CircleStrategy {
	return &CircleStrategy{
		Center: center,
		Radius: model.Radius{
			X: radius,
			Y: radius,
		},
	}
}

func (s *CircleStrategy) Draw(canvas Canvas, color string) {
	canvas.SetColor(color)
	canvas.DrawEllipse(s.Center, s.Radius)
}

func (s *CircleStrategy) MoveShape(vector model.Point) {
	s.Center.X = s.Center.X + vector.X
	s.Center.Y = s.Center.Y + vector.Y
}

func (s *CircleStrategy) GetShapeInfo() string {
	return fmt.Sprintf(
		"%s %.2f %.2f %.2f",
		common.Circle,
		s.Center.X,
		s.Center.Y,
		s.Radius,
	)
}
