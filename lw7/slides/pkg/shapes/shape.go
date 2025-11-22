package shapes

import (
	"slides/pkg/canvas"
	"slides/pkg/model"
)

type Shape interface {
	Draw(canvas canvas.Canvas)
	GetFrame() model.Frame
	GetLineStyle() Style
	GetFillStyle() Style
	SetFrame(model.Frame)
	SetLineStyle(Style)
	SetFillStyle(Style)
	Clone() Shape
}
