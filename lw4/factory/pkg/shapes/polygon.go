package shapes

import (
	"math"

	"factory/pkg/canvas"
	"factory/pkg/model"
)

type polygon struct {
	baseShape
	center      model.Point
	radius      int
	vertexCount int
}

func NewPolygon(color model.Color, center model.Point, radius, vertexCount int) Shape {
	return &polygon{baseShape: baseShape{color: color}, center: center, radius: radius, vertexCount: vertexCount}
}

func (p *polygon) Draw(canvas canvas.Canvas) {
	if p.vertexCount < 3 {
		return
	}
	canvas.SetColor(p.color)

	points := make([]model.Point, p.vertexCount)
	for i := 0; i < p.vertexCount; i++ {
		angle := 2.0 * math.Pi * float64(i) / float64(p.vertexCount)
		points[i] = model.Point{
			X: p.center.X + int(float64(p.radius)*math.Cos(angle)),
			Y: p.center.Y + int(float64(p.radius)*math.Sin(angle)),
		}
	}
	for i := 0; i < p.vertexCount-1; i++ {
		canvas.DrawLine(points[i], points[i+1])
	}
	canvas.DrawLine(points[p.vertexCount-1], points[0])
}
