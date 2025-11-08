package shapedrawinglib

import "adapter/pkg/graphicslib"

type CanvasDrawable interface {
	Draw(canvas graphicslib.Canvas)
}
