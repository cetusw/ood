package shapes

import (
	"math"
	"slides/pkg/canvas"
	"slides/pkg/model"
)

type polygon struct {
	baseShape
	lineStyle Style
	fillStyle Style
	vertices  []model.Point
}

func NewPolygon(vertices []model.Point, lineStyle Style, fillStyle Style) Shape {
	minX, minY := math.MaxFloat64, math.MaxFloat64
	maxX, maxY := -math.MaxFloat64, -math.MaxFloat64
	for _, v := range vertices {
		minX = math.Min(minX, v.X)
		minY = math.Min(minY, v.Y)
		maxX = math.Max(maxX, v.X)
		maxY = math.Max(maxY, v.Y)
	}
	return &polygon{
		baseShape: baseShape{frame: model.Frame{
			Point: model.Point{
				X: minX,
				Y: maxY,
			},
			Width:  maxX - minX,
			Height: maxY - minY,
		}},
		vertices:  vertices,
		lineStyle: lineStyle,
		fillStyle: fillStyle,
	}
}

func (p *polygon) Clone() Shape {
	newVertices := make([]model.Point, len(p.vertices))
	copy(newVertices, p.vertices)

	newPoly := &polygon{
		baseShape: baseShape{
			frame: p.frame,
			color: p.color,
		},
		vertices: newVertices,
	}

	if p.lineStyle != nil {
		newPoly.lineStyle = p.lineStyle.Clone()
	}
	if p.fillStyle != nil {
		newPoly.fillStyle = p.fillStyle.Clone()
	}

	return newPoly
}

func (p *polygon) Draw(c canvas.Canvas) {
	c.SetFillColor(p.fillStyle.GetColor())
	c.SetStrokeColor(p.lineStyle.GetColor())
	c.DrawPolygon(p.vertices)
}

func (p *polygon) GetFrame() model.Frame { return p.frame }

func (p *polygon) SetFrame(newFrame model.Frame) {
	oldFrame := p.frame
	if oldFrame.Width == 0 || oldFrame.Height == 0 {
		return
	}
	scaleX := newFrame.Width / oldFrame.Width
	scaleY := newFrame.Height / oldFrame.Height

	for i := range p.vertices {
		p.vertices[i].X = newFrame.X + (p.vertices[i].X-oldFrame.X)*scaleX
		p.vertices[i].Y = newFrame.Y + (p.vertices[i].Y-oldFrame.Y)*scaleY
	}
	p.frame = newFrame
}

func (p *polygon) GetLineStyle() Style {
	return p.lineStyle
}

func (p *polygon) GetFillStyle() Style {
	return p.fillStyle
}

func (p *polygon) SetLineStyle(s Style) {
	p.lineStyle = s
}

func (p *polygon) SetFillStyle(s Style) {
	p.fillStyle = s
}
