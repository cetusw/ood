package shape

type Shape struct {
	strategy Strategy
	id       string
	color    string
}

func NewShape(strategy Strategy, id string, color string) *Shape {
	return &Shape{
		strategy: strategy,
		id:       id,
		color:    color,
	}
}

func (s *Shape) GetStrategy() Strategy {
	return s.strategy
}

func (s *Shape) GetId() string {
	return s.id
}

func (s *Shape) GetColor() string {
	return s.color
}

func (s *Shape) SetStrategy(strategy Strategy) {
	s.strategy = strategy
}

func (s *Shape) SetId(id string) {
	s.id = id
}

func (s *Shape) SetColor(color string) {
	s.color = color
}

func (s *Shape) Draw(canvas Canvas) string {
	return s.strategy.Draw(canvas, s.id, s.color)
}
