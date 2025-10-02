package tests

import (
	"shapes/pkg/shapes/model"
)

type MockCanvas struct {
	moveToCalled      bool
	setColorValue     string
	lineToCalled      bool
	drawEllipseCalled bool
	saveToFileCalled  bool
	drawText          bool

	lastPoint model.Point
	lastColor string
}

func (m *MockCanvas) MoveTo(p model.Point) {
	m.moveToCalled = true
	m.lastPoint = p
}

func (m *MockCanvas) SetColor(color string) {
	m.setColorValue = color
	m.lastColor = color
}

func (m *MockCanvas) LineTo(p model.Point) {
	m.lineToCalled = true
}

func (m *MockCanvas) DrawEllipse(center model.Point, r model.Radius) {
	m.drawEllipseCalled = true
}

func (m *MockCanvas) DrawText(p model.Point, fontSize float64, text string) {
	m.drawText = true
}

func (m *MockCanvas) SaveToFile(filename string) error {
	m.saveToFileCalled = true
	return nil
}
