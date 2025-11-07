package mocks

import "github.com/stretchr/testify/mock"

type MockItem struct {
	mock.Mock
}

func (m *MockItem) ToHTML() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockItem) ToString() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockItem) GetItemType() string {
	args := m.Called()
	return args.String(0)
}
