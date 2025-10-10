package mocks

import (
	"factory/pkg/canvas"
)

type MockShape struct {
	DrawCalled bool
}

func (m *MockShape) Draw(_ canvas.Canvas) {
	m.DrawCalled = true
}
