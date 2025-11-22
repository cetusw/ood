package shapes

import (
	"slides/pkg/canvas"
	"slides/pkg/model"
	"slides/pkg/style"
)

type Shape interface {
	Draw(canvas canvas.Canvas)
	GetFrame() model.Frame
	GetLineStyle() style.Style
	GetFillStyle() style.Style
	SetFrame(model.Frame)
	SetLineStyle(style.Style)
	SetFillStyle(style.Style)
}
