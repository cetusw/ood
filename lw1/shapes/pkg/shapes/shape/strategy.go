package shape

import "shapes/pkg/shapes/model"

type Strategy interface {
	Draw(canvas Canvas, color string)
	MoveShape(vector model.Point)
	GetShapeInfo() string
}
