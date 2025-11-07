package tests

import (
	"fmt"
	"html"
	"testing"

	"editor/pkg/document"

	"github.com/stretchr/testify/assert"
)

func TestParagraph(t *testing.T) {
	text := "Hello World"
	p := document.NewParagraph(text)

	t.Run("GetText", func(t *testing.T) {
		assert.Equal(t, text, p.GetText())
	})

	t.Run("SetText", func(t *testing.T) {
		newText := "New Text"
		p.SetText(newText)
		assert.Equal(t, newText, p.GetText())
	})

	t.Run("ToHTML", func(t *testing.T) {
		p.SetText(text)
		expectedHTML := fmt.Sprintf("<p>%s</p>", html.EscapeString(text))
		assert.Equal(t, expectedHTML, p.ToHTML())
	})

	t.Run("ToString", func(t *testing.T) {
		expectedString := fmt.Sprintf("Paragraph: %s", text)
		assert.Equal(t, expectedString, p.ToString())
	})

	t.Run("GetParagraph", func(t *testing.T) {
		assert.Equal(t, p, p.GetParagraph())
	})
}
