package shape

type Shape struct {
	strategy Strategy
	shapeID  string
	color    string
}

func NewShape(strategy Strategy, shapeID string, color string) *Shape {
	return &Shape{
		strategy: strategy,
		shapeID:  shapeID,
		color:    color,
	}
}

func (s *Shape) GetStrategy() Strategy {
	return s.strategy
}

func (s *Shape) GetID() string {
	return s.shapeID
}

func (s *Shape) GetColor() string {
	return s.color
}

func (s *Shape) SetStrategy(strategy Strategy) {
	s.strategy = strategy
}

func (s *Shape) SetId(shapeID string) {
	s.shapeID = shapeID
}

func (s *Shape) SetColor(color string) {
	s.color = color
}
