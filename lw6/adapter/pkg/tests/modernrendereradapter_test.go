package tests

import (
	"bytes"
	"strings"
	"testing"

	"adapter/pkg/adapter"
	"adapter/pkg/moderngraphicslib"
)

func TestModernRendererAdapter(t *testing.T) {
	t.Run("single line after move", func(t *testing.T) {
		var buffer bytes.Buffer
		renderer := moderngraphicslib.NewModernGraphicsRenderer(&buffer)
		a := adapter.NewModernRendererAdapter(renderer)

		expectedOutput := strings.Join([]string{
			"<draw>",
			"  <line fromX=\"10\" fromY=\"20\" toX=\"100\" toY=\"150\"/>",
			"</draw>",
			"",
		}, "\n")

		renderer.BeginDraw()
		a.MoveTo(10, 20)
		a.LineTo(100, 150)
		renderer.EndDraw()

		if buffer.String() != expectedOutput {
			t.Errorf("\nExpected:\n%s\nGot:\n%s", expectedOutput, buffer.String())
		}
	})

	t.Run("multiple connected lines", func(t *testing.T) {
		var buffer bytes.Buffer
		renderer := moderngraphicslib.NewModernGraphicsRenderer(&buffer)
		a := adapter.NewModernRendererAdapter(renderer)

		expectedOutput := strings.Join([]string{
			"<draw>",
			"  <line fromX=\"0\" fromY=\"0\" toX=\"10\" toY=\"10\"/>",
			"  <line fromX=\"10\" fromY=\"10\" toX=\"5\" toY=\"15\"/>",
			"</draw>",
			"",
		}, "\n")

		renderer.BeginDraw()
		a.MoveTo(0, 0)
		a.LineTo(10, 10)
		a.LineTo(5, 15)
		renderer.EndDraw()

		if buffer.String() != expectedOutput {
			t.Errorf("\nExpected:\n%s\nGot:\n%s", expectedOutput, buffer.String())
		}
	})

	t.Run("multiple moves should not produce output", func(t *testing.T) {
		var buffer bytes.Buffer
		renderer := moderngraphicslib.NewModernGraphicsRenderer(&buffer)
		a := adapter.NewModernRendererAdapter(renderer)

		expectedOutput := "<draw>\n</draw>\n"

		renderer.BeginDraw()
		a.MoveTo(10, 10)
		a.MoveTo(20, 20)
		renderer.EndDraw()

		if buffer.String() != expectedOutput {
			t.Errorf("\nExpected:\n%s\nGot:\n%s", expectedOutput, buffer.String())
		}
	})
}
