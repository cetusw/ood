package mocks

import "factory/pkg/shapes"

type MockShapeFactory struct {
	ReturnShape         shapes.Shape
	ReturnError         error
	ReceivedDescription string
}

func (m *MockShapeFactory) CreateShape(description string) (shapes.Shape, error) {
	m.ReceivedDescription = description
	if m.ReturnError != nil {
		return nil, m.ReturnError
	}
	return m.ReturnShape, nil
}
