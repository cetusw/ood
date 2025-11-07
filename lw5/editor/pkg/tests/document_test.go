package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"editor/pkg/document"
	"editor/pkg/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createImageFile(t *testing.T, dir, filename string) string {
	t.Helper()
	path := filepath.Join(dir, filename)
	err := os.WriteFile(path, []byte("content"), 0644)
	require.NoError(t, err)
	return path
}

func TestDocument(t *testing.T) {
	t.Run("NewDocument", func(t *testing.T) {
		title := "My Test Document"
		doc := document.NewDocument(title)
		assert.Equal(t, title, doc.GetTitle())
		assert.Equal(t, 0, doc.GetItemsCount())
	})

	t.Run("Set and Get Title", func(t *testing.T) {
		doc := document.NewDocument("Initial Title")
		newTitle := "Updated Title"
		doc.SetTitle(newTitle)
		assert.Equal(t, newTitle, doc.GetTitle())
	})

	t.Run("InsertParagraph", func(t *testing.T) {
		doc := document.NewDocument("Test")
		p1, err := doc.InsertParagraph("p1", 0)
		assert.NoError(t, err)
		assert.Equal(t, 1, doc.GetItemsCount())

		p2, err := doc.InsertParagraph("p2", 1)
		assert.NoError(t, err)
		assert.Equal(t, 2, doc.GetItemsCount())

		p3, err := doc.InsertParagraph("p3", 1)
		assert.NoError(t, err)
		assert.Equal(t, 3, doc.GetItemsCount())

		item, _ := doc.GetItem(0)
		assert.Equal(t, p1, item)

		item, _ = doc.GetItem(1)
		assert.Equal(t, p3, item)

		item, _ = doc.GetItem(2)
		assert.Equal(t, p2, item)
	})

	t.Run("InsertImage", func(t *testing.T) {
		tempDir := t.TempDir()
		imagePath := createImageFile(t, tempDir, "source.png")
		doc := document.NewDocument("Test")

		img, err := doc.InsertImage(imagePath, model.Size{Width: 10, Height: 10}, 0)
		assert.NoError(t, err)
		assert.Equal(t, 1, doc.GetItemsCount())

		item, _ := doc.GetItem(0)
		assert.Equal(t, img, item)

		imgImpl := img.(document.Image)
		assert.FileExists(t, imgImpl.GetPath())
		assert.NotEqual(t, imagePath, imgImpl.GetPath())
	})

	t.Run("InsertItem Paragraph", func(t *testing.T) {
		doc := document.NewDocument("Test")
		p := document.NewParagraph("test")
		item, err := doc.InsertItem(p, 0)
		assert.NoError(t, err)
		assert.Equal(t, p, item)
		assert.Equal(t, 1, doc.GetItemsCount())
	})

	t.Run("InsertItem Image", func(t *testing.T) {
		tempDir := t.TempDir()
		doc := document.NewDocument("Test")
		imagePath := createImageFile(t, tempDir, "source.png")
		img := document.NewImage(model.Size{Width: 10, Height: 10}, imagePath)
		item, err := doc.InsertItem(img, 0)
		assert.NoError(t, err)
		assert.Equal(t, img, item)
		assert.Equal(t, 1, doc.GetItemsCount())
	})

	t.Run("DeleteItem", func(t *testing.T) {
		doc := document.NewDocument("Test")
		_, _ = doc.InsertParagraph("p1", 0)
		_, _ = doc.InsertParagraph("p2", 1)

		deleted, err := doc.DeleteItem(0)
		assert.NoError(t, err)
		assert.Equal(t, 1, doc.GetItemsCount())
		assert.Equal(t, "p1", deleted.(document.Paragraph).GetText())

		item, _ := doc.GetItem(0)
		assert.Equal(t, "p2", item.(document.Paragraph).GetText())

		_, err = doc.DeleteItem(1)
		assert.Error(t, err)
	})

	t.Run("GetItem", func(t *testing.T) {
		doc := document.NewDocument("Test")
		p, _ := doc.InsertParagraph("p1", 0)

		item, err := doc.GetItem(0)
		assert.NoError(t, err)
		assert.Equal(t, p, item)

		_, err = doc.GetItem(1)
		assert.Error(t, err)
	})

	t.Run("Save", func(t *testing.T) {
		tempDir := t.TempDir()
		savePath := filepath.Join(tempDir, "output.html")

		doc := document.NewDocument("Saved Doc")
		_, _ = doc.InsertParagraph("Hello", 0)

		err := doc.Save(savePath)
		assert.NoError(t, err)

		content, err := os.ReadFile(savePath)
		assert.NoError(t, err)

		htmlContent := string(content)
		assert.True(t, strings.Contains(htmlContent, "<title>Saved Doc</title>"))
		assert.True(t, strings.Contains(htmlContent, "<p>Hello</p>"))
	})
}
