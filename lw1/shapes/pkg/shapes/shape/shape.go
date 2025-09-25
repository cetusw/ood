package shape

type Shape struct {
	Strategy Strategy
	Id       string
	Color    string
}

func NewShape(strategy Strategy, id string, color string) *Shape {
	return &Shape{
		Strategy: strategy,
		Id:       id,
		Color:    color,
	}
}

func (s *Shape) SetStrategy(strategy Strategy) {
	s.Strategy = strategy
}

func (s *Shape) Draw(canvas Canvas) string {
	return s.Strategy.Draw(canvas, s.Id, s.Color)
}
