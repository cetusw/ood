package canvas

import (
	"image/color"

	"factory/pkg/domain"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

type PngCanvas struct {
	context *canvas.Context
	c       *canvas.Canvas
}

func NewPngCanvas(width, height float64) *PngCanvas {
	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	ctx.SetFillColor(color.White)
	ctx.DrawPath(0, 0, canvas.Rectangle(width, height))
	ctx.Fill()

	ctx.SetStrokeWidth(2.0)
	return &PngCanvas{
		context: ctx,
		c:       c,
	}
}

func (s *PngCanvas) GetCanvas() *canvas.Canvas {
	return s.c
}

func (s *PngCanvas) SetColor(c domain.Color) {
	s.context.SetStrokeColor(parseColor(c))
}

func (s *PngCanvas) DrawLine(from, to domain.Point) {
	path := &canvas.Path{}
	path.MoveTo(float64(from.X), float64(from.Y))
	path.LineTo(float64(to.X), float64(to.Y))
	s.context.DrawPath(0, 0, path)
	s.context.Stroke()
}

func (s *PngCanvas) DrawEllipse(center domain.Point, hRadius, vRadius int) {
	rx, ry := float64(hRadius), float64(vRadius)
	path := canvas.Ellipse(rx*2, ry*2)
	s.context.DrawPath(float64(center.X)-rx, float64(center.Y)-ry, path)
	s.context.Stroke()
}

func (s *PngCanvas) SaveToFile(filename string) error {
	return renderers.Write(filename, s.c, canvas.DPMM(5.0))
}

func parseColor(c domain.Color) color.RGBA {
	switch c {
	case domain.Red:
		return color.RGBA{R: 211, G: 47, B: 47, A: 255}
	case domain.Green:
		return color.RGBA{R: 76, G: 175, B: 80, A: 255}
	case domain.Blue:
		return color.RGBA{R: 33, G: 150, B: 243, A: 255}
	case domain.Yellow:
		return color.RGBA{R: 255, G: 235, B: 59, A: 255}
	case domain.Pink:
		return color.RGBA{R: 233, G: 30, B: 99, A: 255}
	case domain.Black:
		return color.RGBA{A: 255}
	default:
		return color.RGBA{A: 255}
	}
}
