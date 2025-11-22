package canvas

import (
	"image/color"

	"slides/pkg/model"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

type Canvas interface {
	DrawLine(from model.Point, to model.Point)
	DrawEllipse(center model.Point, radius model.Radius)
	DrawPolygon(vertices []model.Point)
	SaveToFile(filename string) error
	SetStrokeColor(color model.Color)
	SetFillColor(color model.Color)
}

type PngCanvas struct {
	ctx       *canvas.Context
	pngCanvas *canvas.Canvas
}

func NewPngCanvas(width, height float64) *PngCanvas {
	c := canvas.New(width, height)
	ctx := canvas.NewContext(c)

	ctx.SetFillColor(color.White)
	ctx.DrawPath(0, 0, canvas.Rectangle(width, height))
	ctx.Fill()

	ctx.SetStrokeWidth(2.0)
	return &PngCanvas{
		ctx:       ctx,
		pngCanvas: c,
	}
}

func (p *PngCanvas) GetContext() *canvas.Context {
	return p.ctx
}

func (p *PngCanvas) GetCanvas() *canvas.Canvas {
	return p.pngCanvas
}

func (p *PngCanvas) DrawLine(from model.Point, to model.Point) {
	path := &canvas.Path{}
	path.MoveTo(from.X, from.Y)
	path.LineTo(to.X, to.Y)
	p.ctx.DrawPath(0, 0, path)
	p.ctx.Stroke()
}

func (p *PngCanvas) DrawEllipse(center model.Point, radius model.Radius) {
	drawX := center.X - radius.X
	drawY := center.Y - radius.Y

	path := canvas.Ellipse(radius.X*2, radius.Y*2)
	p.ctx.DrawPath(drawX, drawY, path)
	p.ctx.Stroke()
}

func (p *PngCanvas) DrawPolygon(vertices []model.Point) {
	if len(vertices) < 3 {
		return
	}
	path := &canvas.Path{}
	path.MoveTo(vertices[0].X, vertices[0].Y)
	for i := 1; i < len(vertices); i++ {
		path.LineTo(vertices[i].X, vertices[i].Y)
	}
	path.Close()
	p.ctx.DrawPath(0, 0, path)
	p.ctx.Fill()
}

func (p *PngCanvas) SaveToFile(filename string) error {
	return renderers.Write(filename, p.pngCanvas, canvas.DPMM(5.0))
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
		return color.RGBA{A: 0}
	}
}

func (p *PngCanvas) SetStrokeColor(c model.Color) {
	p.ctx.SetStrokeColor(parseColor(c))
}

func (p *PngCanvas) SetFillColor(c model.Color) {
	p.ctx.SetFillColor(parseColor(c))
}
