package shape

import "shapes/pkg/shapes/model"

type Strategy interface {
	Draw(canvas Canvas, id string, color string) string
	MoveShape(vector model.Point)
	GetShapeInfo() string
}
