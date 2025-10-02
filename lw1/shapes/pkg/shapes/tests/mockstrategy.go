package tests

import (
	"shapes/pkg/shapes/model"
	"shapes/pkg/shapes/shape"
)

type MockStrategy struct {
	DrawCalled      bool
	MoveShapeCalled bool
	MoveVector      model.Point
	Info            string
}

func (m *MockStrategy) Draw(canvas shape.Canvas, color string) {
	m.DrawCalled = true
}

func (m *MockStrategy) MoveShape(vector model.Point) {
	m.MoveShapeCalled = true
	m.MoveVector = vector
}

func (m *MockStrategy) GetShapeInfo() string {
	return m.Info
}
