package shape

import "shapes/pkg/shapes/model"

type Canvas interface {
	MoveTo(point model.Point)
	SetColor(color string)
	LineTo(point model.Point)
	DrawEllipse(center model.Point, radius model.Radius)
	DrawPolygon(point model.Point, width float64, height float64)
	DrawText()
}
