package shapes

import (
	"factory/pkg/domain"
	"math"
)

type polygon struct {
	baseShape
	center      domain.Point
	radius      int
	vertexCount int
}

func NewPolygon(color domain.Color, center domain.Point, radius, vertexCount int) domain.Shape {
	return &polygon{baseShape: baseShape{color: color}, center: center, radius: radius, vertexCount: vertexCount}
}

func (p *polygon) Draw(canvas domain.Canvas) {
	if p.vertexCount < 3 {
		return
	}
	canvas.SetColor(p.color)

	points := make([]domain.Point, p.vertexCount)
	for i := 0; i < p.vertexCount; i++ {
		angle := 2.0 * math.Pi * float64(i) / float64(p.vertexCount)
		points[i] = domain.Point{
			X: p.center.X + int(float64(p.radius)*math.Cos(angle)),
			Y: p.center.Y + int(float64(p.radius)*math.Sin(angle)),
		}
	}
	for i := 0; i < p.vertexCount-1; i++ {
		canvas.DrawLine(points[i], points[i+1])
	}
	canvas.DrawLine(points[p.vertexCount-1], points[0])
}
