package shape

import (
	"fmt"
	"shapes/pkg/common"
	"shapes/pkg/shapes/model"
)

type CircleStrategy struct {
	Center model.Point
	Radius float64
}

func NewCircleStrategy(center model.Point, radius float64) *CircleStrategy {
	return &CircleStrategy{
		Center: center,
		Radius: radius,
	}
}

func (s *CircleStrategy) Draw(canvas Canvas, id string, color string) string {
	return fmt.Sprintf(
		"circle drawn: id: <%s>, color <%s>, Center <%.2f, %.2f>, Radius <%.2f>",
		id,
		color,
		s.Center.X,
		s.Center.Y,
		s.Radius,
	)
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
