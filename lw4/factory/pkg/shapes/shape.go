package shapes

import "factory/pkg/canvas"

type Shape interface {
	Draw(canvas canvas.Canvas)
}
