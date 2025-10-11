package canvas

import (
	"image/color"

	"factory/pkg/model"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

type Canvas interface {
	SetColor(color model.Color)
	DrawLine(from model.Point, to model.Point)
	DrawEllipse(center model.Point, radius model.Radius)
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

func (p *PngCanvas) GetCanvas() *canvas.Canvas {
	return p.C
}

func (p *PngCanvas) SetColor(c model.Color) {
	p.Context.SetStrokeColor(parseColor(c))
}

func (p *PngCanvas) DrawLine(from model.Point, to model.Point) {
	path := &canvas.Path{}
	path.MoveTo(from.X, from.Y)
	path.LineTo(to.X, to.Y)
	p.Context.DrawPath(0, 0, path)
	p.Context.Stroke()
}

func (p *PngCanvas) DrawEllipse(center model.Point, radius model.Radius) {
	path := canvas.Ellipse(radius.X*2, radius.Y*2)
	p.Context.DrawPath(center.X-radius.X, center.Y-radius.Y, path)
	p.Context.Stroke()
}

func (p *PngCanvas) SaveToFile(filename string) error {
	return renderers.Write(filename, p.C, canvas.DPMM(5.0))
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
