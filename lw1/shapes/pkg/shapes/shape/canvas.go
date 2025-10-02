package shape

import (
	"image/color"
	"shapes/pkg/common"
	"shapes/pkg/shapes/model"

	tdw "github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers"
)

type Canvas interface {
	MoveTo(point model.Point)
	SetColor(color string)
	LineTo(point model.Point)
	DrawEllipse(center model.Point, radius model.Radius)
	DrawText(topLeftPoint model.Point, fontSize float64, text string)
	SaveToFile(filename string) error
}

type TDWCanvas struct {
	canvas       *tdw.Canvas
	context      *tdw.Context
	width        float64
	height       float64
	currentPoint model.Point
	currentColor color.RGBA
}

func NewCanvas(width float64, height float64) Canvas {
	canvas := tdw.New(width, height)
	context := tdw.NewContext(canvas)

	bgColor := color.RGBA{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}

	context.SetFillColor(bgColor)
	context.DrawPath(0, 0, tdw.Rectangle(width, height))

	initialColor := color.RGBA{R: 0, G: 0, B: 0, A: 255}

	return &TDWCanvas{
		canvas:       canvas,
		context:      context,
		width:        width,
		height:       height,
		currentPoint: model.Point{X: 0, Y: 0},
		currentColor: initialColor,
	}
}

func (c *TDWCanvas) MoveTo(point model.Point) {
	c.currentPoint = point
}

func (c *TDWCanvas) SetColor(colorStr string) {
	parsedColor, err := common.ParseHexColor(colorStr)
	if err != nil {
		return
	}
	c.currentColor = parsedColor
}

func (c *TDWCanvas) LineTo(point model.Point) {
	currentPath := &tdw.Path{}
	currentPath.MoveTo(c.currentPoint.X, c.currentPoint.Y)
	currentPath.LineTo(point.X, point.Y)

	c.context.SetStrokeColor(c.currentColor)
	c.context.SetStrokeWidth(1.0)
	c.context.DrawPath(0, 0, currentPath)

	c.currentPoint = point
}

func (c *TDWCanvas) DrawEllipse(center model.Point, radius model.Radius) {
	currentPath := &tdw.Path{}
	currentPath.MoveTo(center.X+radius.X, center.Y)
	currentPath.Arc(radius.X, radius.Y, 0, 0, 360)

	c.context.SetStrokeColor(c.currentColor)
	c.context.SetStrokeWidth(1.0)
	c.context.DrawPath(0, 0, currentPath)
}

func (c *TDWCanvas) DrawText(topLeftPoint model.Point, fontSize float64, text string) {
	face := tdw.NewFontFamily("text-font")
	face.LoadSystemFont("Arial", tdw.FontRegular)
	font := face.Face(fontSize, c.currentColor, tdw.FontRegular, tdw.FontNormal)
	c.context.DrawText(topLeftPoint.X, topLeftPoint.Y, tdw.NewTextLine(font, text, tdw.Left))
}

func (c *TDWCanvas) SaveToFile(filename string) error {
	return renderers.Write(filename, c.canvas, tdw.DPMM(5.0))
}
