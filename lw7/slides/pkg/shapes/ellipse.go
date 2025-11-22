package shapes

import (
	"slides/pkg/canvas"
	"slides/pkg/model"
)

type ellipse struct {
	baseShape
	lineStyle Style
	fillStyle Style
	center    model.Point
	radius    model.Radius
}

func NewEllipse(center model.Point, radius model.Radius) Shape {
	left := center.X - radius.X
	bottom := center.Y - radius.Y

	width := radius.X * 2
	height := radius.Y * 2

	e := &ellipse{
		baseShape: baseShape{
			frame: model.Frame{
				Point: model.Point{
					X: left,
					Y: bottom,
				},
				Width:  width,
				Height: height,
			},
		},
		center: center,
		radius: radius,
	}
	return e
}

func (e *ellipse) Clone() Shape {
	newEllipse := &ellipse{
		baseShape: baseShape{
			frame: e.frame,
			color: e.color,
		},
		center: e.center,
		radius: e.radius,
	}

	if e.lineStyle != nil {
		newEllipse.lineStyle = e.lineStyle.Clone()
	}
	if e.fillStyle != nil {
		newEllipse.fillStyle = e.fillStyle.Clone()
	}

	return newEllipse
}

func (e *ellipse) Draw(canvas canvas.Canvas) {
	radius := model.Radius{
		X: e.frame.Width / 2,
		Y: e.frame.Height / 2,
	}
	center := model.Point{
		X: e.frame.X + radius.X,
		Y: e.frame.Y + radius.Y,
	}

	canvas.SetFillColor(e.fillStyle.GetColor())
	canvas.SetStrokeColor(e.lineStyle.GetColor())
	canvas.DrawEllipse(center, radius)
}

func (e *ellipse) GetFrame() model.Frame { return e.frame }

func (e *ellipse) SetFrame(frame model.Frame) {
	e.frame = frame
}

func (e *ellipse) GetLineStyle() Style {
	return e.lineStyle
}
func (e *ellipse) GetFillStyle() Style {
	return e.fillStyle
}

func (e *ellipse) SetLineStyle(s Style) {
	e.lineStyle = s
}
func (e *ellipse) SetFillStyle(s Style) {
	e.fillStyle = s
}
