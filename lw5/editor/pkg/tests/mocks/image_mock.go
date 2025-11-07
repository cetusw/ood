package mocks

import (
	"editor/pkg/document"
	"editor/pkg/model"

	"github.com/stretchr/testify/mock"
)

type MockImage struct {
	MockItem
	mock.Mock
}

func (m *MockImage) Resize(size model.Size) {
	m.Called(size)
}

func (m *MockImage) Destroy() {
	m.Called()
}

func (m *MockImage) GetPath() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockImage) GetSize() model.Size {
	args := m.Called()
	return args.Get(0).(model.Size)
}

func (m *MockImage) GetImage() document.Image {
	args := m.Called()
	return args.Get(0).(document.Image)
}
