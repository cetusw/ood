package canvas

import (
	"image/color"

	"factory/pkg/model"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

type Canvas interface {
	SetColor(color model.Color)
	DrawLine(from, to model.Point)
	DrawEllipse(center model.Point, hRadius, vRadius int)
	SaveToFile(filename string) error
}

type PngCanvas struct {
	Context *canvas.Context
	C       *canvas.Canvas
}

func NewPngCanvas(width, height float64) *PngCanvas {
	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	ctx.SetFillColor(color.White)
	ctx.DrawPath(0, 0, canvas.Rectangle(width, height))
	ctx.Fill()

	ctx.SetStrokeWidth(2.0)
	return &PngCanvas{
		Context: ctx,
		C:       c,
	}
}

func (s *PngCanvas) GetCanvas() *canvas.Canvas {
	return s.C
}

func (s *PngCanvas) SetColor(c model.Color) {
	s.Context.SetStrokeColor(parseColor(c))
}

func (s *PngCanvas) DrawLine(from, to model.Point) {
	path := &canvas.Path{}
	path.MoveTo(float64(from.X), float64(from.Y))
	path.LineTo(float64(to.X), float64(to.Y))
	s.Context.DrawPath(0, 0, path)
	s.Context.Stroke()
}

func (s *PngCanvas) DrawEllipse(center model.Point, hRadius, vRadius int) {
	rx, ry := float64(hRadius), float64(vRadius)
	path := canvas.Ellipse(rx*2, ry*2)
	s.Context.DrawPath(float64(center.X)-rx, float64(center.Y)-ry, path)
	s.Context.Stroke()
}

func (s *PngCanvas) SaveToFile(filename string) error {
	return renderers.Write(filename, s.C, canvas.DPMM(5.0))
}

func parseColor(c model.Color) color.RGBA {
	switch c {
	case model.Red:
		return color.RGBA{R: 211, G: 47, B: 47, A: 255}
	case model.Green:
		return color.RGBA{R: 76, G: 175, B: 80, A: 255}
	case model.Blue:
		return color.RGBA{R: 33, G: 150, B: 243, A: 255}
	case model.Yellow:
		return color.RGBA{R: 255, G: 235, B: 59, A: 255}
	case model.Pink:
		return color.RGBA{R: 233, G: 30, B: 99, A: 255}
	case model.Black:
		return color.RGBA{A: 255}
	default:
		return color.RGBA{A: 255}
	}
}
