package mocks

import (
	"editor/pkg/command"

	"github.com/stretchr/testify/mock"
)

type MockCommand struct {
	mock.Mock
}

func (m *MockCommand) Execute() {
	m.Called()
}

func (m *MockCommand) Unexecute() {
	m.Called()
}

func (m *MockCommand) Destroy() {
	m.Called()
}

func (m *MockCommand) Merge(next command.Command) bool {
	args := m.Called(next)
	return args.Bool(0)
}
