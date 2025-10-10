package domain

type Canvas interface {
	SetColor(color Color)
	DrawLine(from, to Point)
	DrawEllipse(center Point, hRadius, vRadius int)
	SaveToFile(filename string) error
}

type Shape interface {
	Draw(canvas Canvas)
}

type PictureDraft struct {
	shapes []Shape
}

func (p *PictureDraft) AddShape(s Shape) {
	p.shapes = append(p.shapes, s)
}

func (p *PictureDraft) Draw(canvas Canvas) {
	for _, shape := range p.shapes {
		shape.Draw(canvas)
	}
}

type Color string

const (
	Green  Color = "green"
	Red    Color = "red"
	Blue   Color = "blue"
	Yellow Color = "yellow"
	Pink   Color = "pink"
	Black  Color = "black"
)

type Point struct {
	X, Y int
}
