package shapedrawinglib

import (
	"adapter/pkg/graphicslib"
)

type CanvasPainter interface {
	Draw(drawable CanvasDrawable)
}

type canvasPainter struct {
	canvas graphicslib.Canvas
}

func NewCanvasPainter(canvas graphicslib.Canvas) CanvasPainter {
	return &canvasPainter{canvas}
}

func (p *canvasPainter) Draw(drawable CanvasDrawable) {
	drawable.Draw(p.canvas)
}
