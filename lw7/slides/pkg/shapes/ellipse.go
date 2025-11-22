package shapes

import (
	"slides/pkg/canvas"
	"slides/pkg/model"
	"slides/pkg/style"
)

type ellipse struct {
	baseShape
	lineStyle style.Style
	fillStyle style.Style
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

func (e *ellipse) GetLineStyle() style.Style {
	return e.lineStyle
}
func (e *ellipse) GetFillStyle() style.Style {
	return e.fillStyle
}

func (e *ellipse) SetLineStyle(s style.Style) {
	e.lineStyle = s
}
func (e *ellipse) SetFillStyle(s style.Style) {
	e.fillStyle = s
}
