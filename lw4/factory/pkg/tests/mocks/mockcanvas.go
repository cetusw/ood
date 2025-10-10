package mocks

import (
	"fmt"
	"strings"

	"factory/pkg/model"
)

type MockCanvas struct {
	calls []string
}

func (m *MockCanvas) SetColor(color model.Color) {
	m.calls = append(m.calls, fmt.Sprintf("SetColor(%s)", color))
}

func (m *MockCanvas) DrawLine(from, to model.Point) {
	m.calls = append(m.calls, fmt.Sprintf("DrawLine(from: %v, to: %v)", from, to))
}

func (m *MockCanvas) DrawEllipse(center model.Point, hRadius, vRadius int) {
	m.calls = append(m.calls, fmt.Sprintf("DrawEllipse(center: %v, h: %d, v: %d)", center, hRadius, vRadius))
}

func (m *MockCanvas) SaveToFile(filename string) error {
	m.calls = append(m.calls, fmt.Sprintf("SaveToFile(%s)", filename))
	return nil
}

func (m *MockCanvas) GetLog() string {
	return strings.Join(m.calls, "; ")
}
