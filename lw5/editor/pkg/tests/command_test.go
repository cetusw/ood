package tests

import (
	"editor/pkg/model"
	"os"
	"path/filepath"
	"testing"

	"editor/pkg/command"
	"editor/pkg/tests/mocks"

	"github.com/stretchr/testify/assert"
)

func TestInsertParagraphCommand(t *testing.T) {
	mockDoc := new(mocks.MockDocument)
	mockItem := new(mocks.MockItem)
	text := "hello"
	pos := 0

	cmd := command.NewInsertParagraphCommand(mockDoc, pos, text)

	t.Run("Execute", func(t *testing.T) {
		mockDoc.On("InsertParagraph", text, pos).Return(mockItem, nil).Once()
		cmd.Execute()
		mockDoc.AssertExpectations(t)
	})

	t.Run("Unexecute", func(t *testing.T) {
		mockDoc.On("DeleteItem", pos).Return(mockItem, nil).Once()
		cmd.Unexecute()
		mockDoc.AssertExpectations(t)
	})

	t.Run("Merge", func(t *testing.T) {
		assert.False(t, cmd.Merge(nil))
	})
}

func TestDeleteItemCommand(t *testing.T) {
	mockDoc := new(mocks.MockDocument)
	mockParagraph := new(mocks.MockParagraph)
	mockImage := new(mocks.MockImage)
	pos := 0

	t.Run("Execute and Unexecute paragraph", func(t *testing.T) {
		cmd := command.NewDeleteItemCommand(mockDoc, pos)

		mockDoc.On("DeleteItem", pos).Return(mockParagraph, nil).Once()
		cmd.Execute()
		mockDoc.AssertExpectations(t)

		mockDoc.On("InsertItem", mockParagraph, pos).Return(mockParagraph, nil).Once()
		cmd.Unexecute()
		mockDoc.AssertExpectations(t)
	})

	t.Run("Destroy with image", func(t *testing.T) {
		cmd := command.NewDeleteItemCommand(mockDoc, pos)

		mockDoc.On("DeleteItem", pos).Return(mockImage, nil).Once()
		cmd.Execute()

		mockImage.On("Destroy").Return().Once()
		cmd.Destroy()
		mockImage.AssertExpectations(t)
	})

	t.Run("Destroy with paragraph", func(t *testing.T) {
		cmd := command.NewDeleteItemCommand(mockDoc, pos)

		mockDoc.On("DeleteItem", pos).Return(mockParagraph, nil).Once()
		cmd.Execute()

		cmd.Destroy()
	})
}

func TestSetTitleCommand(t *testing.T) {
	mockDoc := new(mocks.MockDocument)
	oldTitle := "Old Title"
	newTitle := "New Title"

	cmd := command.NewSetTitleCommand(mockDoc, newTitle)

	t.Run("Execute", func(t *testing.T) {
		mockDoc.On("GetTitle").Return(oldTitle).Once()
		mockDoc.On("SetTitle", newTitle).Return().Once()
		cmd.Execute()
		mockDoc.AssertExpectations(t)
	})

	t.Run("Unexecute", func(t *testing.T) {
		mockDoc.On("SetTitle", oldTitle).Return().Once()
		cmd.Unexecute()
		mockDoc.AssertExpectations(t)
	})

	t.Run("Merge", func(t *testing.T) {
		anotherTitle := "Another Title"
		nextCmd := command.NewSetTitleCommand(mockDoc, anotherTitle)

		canMerge := cmd.Merge(nextCmd)
		assert.True(t, canMerge)

		assert.False(t, cmd.Merge(nil), "Should not merge with non-setTitle command")
	})
}

func TestInsertImageCommand(t *testing.T) {
	mockDoc := new(mocks.MockDocument)
	mockItem := new(mocks.MockItem)
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "test.jpg")
	err := os.WriteFile(path, []byte("content"), 0644)
	assert.NoError(t, err)
	size := model.Size{Width: 100, Height: 200}
	pos := 1

	cmd := command.NewInsertImageCommand(mockDoc, path, size, pos)

	t.Run("Execute", func(t *testing.T) {
		mockDoc.On("InsertImage", path, size, pos).Return(mockItem, nil).Once()
		cmd.Execute()
		mockDoc.AssertExpectations(t)
	})

	t.Run("Unexecute", func(t *testing.T) {
		mockDoc.On("DeleteItem", pos).Return(mockItem, nil).Once()
		cmd.Unexecute()
		mockDoc.AssertExpectations(t)
	})

	t.Run("Destroy with inserted item", func(t *testing.T) {
		mockDocDestroy := new(mocks.MockDocument)
		mockImageDestroy := new(mocks.MockImage)

		cmdDestroy := command.NewInsertImageCommand(mockDocDestroy, path, size, pos)

		mockDocDestroy.On("InsertImage", path, size, pos).Return(mockImageDestroy, nil).Once()
		cmdDestroy.Execute()

		mockImageDestroy.On("Destroy").Return().Once()
		cmdDestroy.Destroy()
		mockImageDestroy.AssertExpectations(t)
	})

	t.Run("Merge", func(t *testing.T) {
		assert.False(t, cmd.Merge(nil))
	})
}

func TestReplaceTextCommand(t *testing.T) {
	mockDoc := new(mocks.MockDocument)
	mockParagraph := new(mocks.MockParagraph)
	pos := 0
	oldText := "Старый текст"
	newText := "Новый текст"

	cmd := command.NewReplaceTextCommand(mockDoc, pos, newText)

	t.Run("Execute", func(t *testing.T) {
		mockDoc.On("GetItem", pos).Return(mockParagraph, nil).Once()
		mockParagraph.On("GetParagraph").Return(mockParagraph).Once()
		mockParagraph.On("GetText").Return(oldText).Once()
		mockParagraph.On("SetText", newText).Return().Once()

		cmd.Execute()

		mockDoc.AssertExpectations(t)
		mockParagraph.AssertExpectations(t)
	})

	t.Run("Unexecute", func(t *testing.T) {
		mockParagraph.On("SetText", oldText).Return().Once()

		cmd.Unexecute()

		mockParagraph.AssertExpectations(t)
	})

	t.Run("Merge", func(t *testing.T) {
		anotherText := "Совсем новый текст"
		nextCmd := command.NewReplaceTextCommand(mockDoc, pos, anotherText)

		canMerge := cmd.Merge(nextCmd)
		assert.True(t, canMerge)

		mockParagraph.On("SetText", oldText).Return().Once()
		cmd.Unexecute()
		mockParagraph.AssertExpectations(t)
	})
}

func TestResizeImageCommand(t *testing.T) {
	mockDoc := new(mocks.MockDocument)
	mockImage := new(mocks.MockImage)
	pos := 0
	oldSize := model.Size{Width: 100, Height: 100}
	newSize := model.Size{Width: 200, Height: 250}

	cmd := command.NewResizeImageCommand(mockDoc, pos, newSize)

	t.Run("Execute", func(t *testing.T) {
		mockDoc.On("GetItem", pos).Return(mockImage, nil).Once()
		mockImage.On("GetImage").Return(mockImage).Once()
		mockImage.On("GetSize").Return(oldSize).Once()
		mockImage.On("Resize", newSize).Return().Once()

		cmd.Execute()

		mockDoc.AssertExpectations(t)
		mockImage.AssertExpectations(t)
	})

	t.Run("Unexecute", func(t *testing.T) {
		mockImage.On("Resize", oldSize).Return().Once()

		cmd.Unexecute()

		mockImage.AssertExpectations(t)
	})

	t.Run("Merge", func(t *testing.T) {
		anotherSize := model.Size{Width: 50, Height: 50}
		nextCmd := command.NewResizeImageCommand(mockDoc, pos, anotherSize)

		canMerge := cmd.Merge(nextCmd)
		assert.True(t, canMerge)

		mockImage.On("Resize", oldSize).Return().Once()
		cmd.Unexecute()
		mockImage.AssertExpectations(t)
	})
}
