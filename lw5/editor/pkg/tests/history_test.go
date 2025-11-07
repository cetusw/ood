package tests

import (
	"editor/pkg/history"
	"editor/pkg/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHistory(t *testing.T) {
	t.Run("AddAndExecuteCommand", func(t *testing.T) {
		h := history.NewHistory()
		cmd := new(mocks.MockCommand)

		cmd.On("Execute").Return().Once()

		h.AddAndExecuteCommand(cmd)

		cmd.AssertExpectations(t)
		assert.True(t, h.CanUndo())
		assert.False(t, h.CanRedo())
	})

	t.Run("Undo and Redo", func(t *testing.T) {
		h := history.NewHistory()
		cmd := new(mocks.MockCommand)

		cmd.On("Execute").Return().Times(2)
		cmd.On("Unexecute").Return().Once()

		h.AddAndExecuteCommand(cmd)
		h.Undo()
		h.Redo()

		cmd.AssertExpectations(t)
		assert.True(t, h.CanUndo())
		assert.False(t, h.CanRedo())
	})

	t.Run("Merge commands", func(t *testing.T) {
		h := history.NewHistory()
		cmd1 := new(mocks.MockCommand)
		cmd2 := new(mocks.MockCommand)

		cmd1.On("Execute").Return().Once()
		cmd1.On("Merge", cmd2).Return(true).Once()
		cmd1.On("Unexecute").Return().Once()
		cmd2.On("Execute").Return().Once()

		h.AddAndExecuteCommand(cmd1)
		h.AddAndExecuteCommand(cmd2)

		h.Undo()

		cmd1.AssertExpectations(t)
		cmd2.AssertExpectations(t)

		assert.False(t, h.CanUndo())
	})

	t.Run("Adding command after undo truncates redo stack", func(t *testing.T) {
		h := history.NewHistory()
		cmd1 := new(mocks.MockCommand)
		cmd2 := new(mocks.MockCommand)
		cmd3 := new(mocks.MockCommand)

		cmd1.On("Merge", cmd2).Return(false).Once()
		cmd1.On("Merge", cmd3).Return(false).Once()
		cmd1.On("Execute").Return().Once()
		cmd1.On("Unexecute").Return().Once()

		cmd2.On("Execute").Return().Once()
		cmd2.On("Unexecute").Return().Once()
		cmd2.On("Destroy").Return().Once()

		cmd3.On("Execute").Return().Once()
		cmd3.On("Unexecute").Return().Once()

		h.AddAndExecuteCommand(cmd1)
		h.AddAndExecuteCommand(cmd2)

		h.Undo()
		assert.True(t, h.CanRedo())

		h.AddAndExecuteCommand(cmd3)
		assert.False(t, h.CanRedo(), "Redo stack should be cleared")

		h.Undo()
		h.Undo()

		cmd1.AssertExpectations(t)
		cmd2.AssertExpectations(t)
		cmd3.AssertExpectations(t)
	})
}
