package mocks

import (
	"editor/pkg/document"

	"github.com/stretchr/testify/mock"
)

type MockParagraph struct {
	MockItem
	mock.Mock
}

func (m *MockParagraph) GetText() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockParagraph) SetText(text string) {
	m.Called(text)
}

func (m *MockParagraph) GetParagraph() document.Paragraph {
	args := m.Called()
	return args.Get(0).(document.Paragraph)
}
