package tests

import (
	"fmt"
	"html"
	"os"
	"path/filepath"
	"testing"

	"editor/pkg/document"
	"editor/pkg/model"

	"github.com/stretchr/testify/assert"
)

func TestImage(t *testing.T) {
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "test.jpg")
	err := os.WriteFile(tempFile, []byte("content"), 0644)
	assert.NoError(t, err)

	size := model.Size{Width: 100, Height: 200}
	img := document.NewImage(size, tempFile)

	t.Run("Getters", func(t *testing.T) {
		assert.Equal(t, tempFile, img.GetPath())
		assert.Equal(t, size, img.GetSize())
		assert.Equal(t, img, img.GetImage())
	})

	t.Run("ToHTML", func(t *testing.T) {
		expectedHTML := fmt.Sprintf(
			`<img src="%s" width="%d" height="%d" />`,
			html.EscapeString(tempFile), size.Width, size.Height,
		)
		assert.Equal(t, expectedHTML, img.ToHTML())
	})

	t.Run("ToString", func(t *testing.T) {
		expectedString := fmt.Sprintf("Image: %d %d %s", size.Width, size.Height, tempFile)
		assert.Equal(t, expectedString, img.ToString())
	})

	t.Run("Resize", func(t *testing.T) {
		newSize := model.Size{Width: 300, Height: 400}
		img.Resize(newSize)
		assert.Equal(t, newSize, img.GetSize())
	})

	t.Run("Destroy", func(t *testing.T) {
		destroyableFilePath := filepath.Join(tempDir, "destroy.jpg")
		err := os.WriteFile(destroyableFilePath, []byte("content"), 0644)
		assert.NoError(t, err)

		destroyableImage := document.NewImage(size, destroyableFilePath)

		_, err = os.Stat(destroyableFilePath)
		assert.NoError(t, err)

		destroyableImage.Destroy()

		_, err = os.Stat(destroyableFilePath)
		assert.True(t, os.IsNotExist(err))
	})
}
