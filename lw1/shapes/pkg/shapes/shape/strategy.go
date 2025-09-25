package shape

type Strategy interface {
	Draw(canvas Canvas, id string, color string) string
	GetShapeInfo() string
}
