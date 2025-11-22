package shapes

import (
	"testing"

	"slides/pkg/canvas"
	"slides/pkg/model"
	"slides/pkg/style"

	"github.com/stretchr/testify/assert"
)

type mockShape struct {
	frame     model.Frame
	lineStyle style.Style
	fillStyle style.Style
}

func newMockShape(f model.Frame, ls, fs style.Style) *mockShape {
	return &mockShape{
		frame:     f,
		lineStyle: ls,
		fillStyle: fs,
	}
}

func (m *mockShape) Draw(_ canvas.Canvas) {}

func (m *mockShape) GetFrame() model.Frame  { return m.frame }
func (m *mockShape) SetFrame(f model.Frame) { m.frame = f }

func (m *mockShape) GetLineStyle() style.Style  { return m.lineStyle }
func (m *mockShape) SetLineStyle(s style.Style) { m.lineStyle = s }

func (m *mockShape) GetFillStyle() style.Style  { return m.fillStyle }
func (m *mockShape) SetFillStyle(s style.Style) { m.fillStyle = s }

func TestGroup_EmptyGroup(t *testing.T) {
	group := NewGroup()
	f := group.GetFrame()

	assert.Equal(t, 0.0, f.Width)
	assert.Equal(t, 0.0, f.Height)
}

func TestGroup_NonIntersectingShapes(t *testing.T) {
	s1 := newMockShape(
		model.Frame{Point: model.Point{X: 0, Y: 0}, Width: 10, Height: 10},
		nil, nil,
	)
	s2 := newMockShape(
		model.Frame{Point: model.Point{X: 20, Y: 20}, Width: 10, Height: 10},
		nil, nil,
	)

	group := NewGroup()
	group.AddShape(s1)
	group.AddShape(s2)

	assert.Equal(t, 0.0, group.GetFrame().X)
	assert.Equal(t, 20.0, group.GetFrame().Y)
	assert.Equal(t, 30.0, group.GetFrame().Width)
	assert.Equal(t, 30.0, group.GetFrame().Height)
}

func TestGroup_IntersectingShapes(t *testing.T) {
	s1 := newMockShape(
		model.Frame{Point: model.Point{X: 10, Y: 30}, Width: 20, Height: 20},
		nil, nil,
	)
	s2 := newMockShape(
		model.Frame{Point: model.Point{X: 20, Y: 40}, Width: 20, Height: 20},
		nil, nil,
	)

	group := NewGroup()
	group.AddShape(s1)
	group.AddShape(s2)

	assert.Equal(t, 10.0, group.GetFrame().X)
	assert.Equal(t, 40.0, group.GetFrame().Y)
	assert.Equal(t, 30.0, group.GetFrame().Width)
	assert.Equal(t, 30.0, group.GetFrame().Height)
}

func TestGroup_ResizesGroupFrameWithOneShape(t *testing.T) {
	s := newMockShape(
		model.Frame{Point: model.Point{X: 10, Y: 10}, Width: 10, Height: 10},
		nil, nil,
	)
	group := NewGroup()
	group.AddShape(s)

	assert.Equal(t, 10.0, group.GetFrame().X)
	assert.Equal(t, 10.0, group.GetFrame().Y)
	assert.Equal(t, 10.0, group.GetFrame().Width)
	assert.Equal(t, 10.0, group.GetFrame().Height)

	newFrame := model.Frame{
		Point:  model.Point{X: 50, Y: 50},
		Width:  200,
		Height: 200,
	}
	group.SetFrame(newFrame)

	assert.Equal(t, 50.0, s.GetFrame().X)
	assert.Equal(t, 50.0, s.GetFrame().Y)
	assert.Equal(t, 200.0, s.GetFrame().Width)
	assert.Equal(t, 200.0, s.GetFrame().Height)
}

func TestGroup_ResizesGroupFrameWithNonIntersectionShapes(t *testing.T) {
	s1 := newMockShape(
		model.Frame{Point: model.Point{X: 0, Y: 0}, Width: 10, Height: 10},
		nil, nil,
	)
	s2 := newMockShape(
		model.Frame{Point: model.Point{X: 20, Y: 20}, Width: 10, Height: 10},
		nil, nil,
	)

	group := NewGroup()
	group.AddShape(s1)
	group.AddShape(s2)

	assert.Equal(t, 0.0, group.GetFrame().X)
	assert.Equal(t, 20.0, group.GetFrame().Y)
	assert.Equal(t, 30.0, group.GetFrame().Width)
	assert.Equal(t, 30.0, group.GetFrame().Height)

	newFrame := model.Frame{
		Point:  model.Point{X: 0, Y: 20},
		Width:  60,
		Height: 60,
	}
	group.SetFrame(newFrame)

	assert.Equal(t, 0.0, s1.GetFrame().X)
	assert.Equal(t, -20.0, s1.GetFrame().Y)
	assert.Equal(t, 20.0, s1.GetFrame().Width)
	assert.Equal(t, 20.0, s1.GetFrame().Height)

	assert.Equal(t, 40.0, s2.GetFrame().X)
	assert.Equal(t, 20.0, s2.GetFrame().Y)
	assert.Equal(t, 20.0, s2.GetFrame().Width)
	assert.Equal(t, 20.0, s2.GetFrame().Height)
}

func TestGroup_ResizesGroupFrameWithIntersectionShapes(t *testing.T) {
	s1 := newMockShape(
		model.Frame{Point: model.Point{X: 10, Y: 30}, Width: 20, Height: 20},
		nil, nil,
	)
	s2 := newMockShape(
		model.Frame{Point: model.Point{X: 20, Y: 40}, Width: 20, Height: 20},
		nil, nil,
	)

	group := NewGroup()
	group.AddShape(s1)
	group.AddShape(s2)

	assert.Equal(t, 10.0, group.GetFrame().X)
	assert.Equal(t, 40.0, group.GetFrame().Y)
	assert.Equal(t, 30.0, group.GetFrame().Width)
	assert.Equal(t, 30.0, group.GetFrame().Height)

	newFrame := model.Frame{
		Point:  model.Point{X: 10, Y: 40},
		Width:  60,
		Height: 60,
	}
	group.SetFrame(newFrame)

	assert.Equal(t, 10.0, s1.GetFrame().X)
	assert.Equal(t, 20.0, s1.GetFrame().Y)
	assert.Equal(t, 40.0, s1.GetFrame().Width)
	assert.Equal(t, 40.0, s1.GetFrame().Height)

	assert.Equal(t, 30.0, s2.GetFrame().X)
	assert.Equal(t, 40.0, s2.GetFrame().Y)
	assert.Equal(t, 40.0, s2.GetFrame().Width)
	assert.Equal(t, 40.0, s2.GetFrame().Height)
}

func TestGroup_ReturnsCommonLineStyle(t *testing.T) {
	s1 := newMockShape(model.Frame{}, style.NewStyle(true, model.Red), nil)
	s2 := newMockShape(model.Frame{}, style.NewStyle(true, model.Red), nil)

	g := NewGroup()
	g.AddShape(s1)
	g.AddShape(s2)

	assert.Equal(t, model.Red, g.GetLineStyle().GetColor())
	assert.True(t, g.GetLineStyle().IsEnabled())
}

func TestGroup_ReturnsCommonFillStyle(t *testing.T) {
	s1 := newMockShape(model.Frame{}, nil, style.NewStyle(true, model.Red))
	s2 := newMockShape(model.Frame{}, nil, style.NewStyle(true, model.Red))

	g := NewGroup()
	g.AddShape(s1)
	g.AddShape(s2)

	assert.Equal(t, model.Red, g.GetFillStyle().GetColor())
	assert.True(t, g.GetFillStyle().IsEnabled())
}

func TestGroup_ReturnsUndefinedLineStyle(t *testing.T) {
	s1 := newMockShape(model.Frame{}, style.NewStyle(true, model.Red), nil)
	s2 := newMockShape(model.Frame{}, style.NewStyle(true, model.Blue), nil)

	g := NewGroup()
	g.AddShape(s1)
	g.AddShape(s2)

	assert.Equal(t, model.Undefined, g.GetLineStyle().GetColor())
	assert.True(t, g.GetLineStyle().IsEnabled())
}

func TestGroup_ReturnsUndefinedFillStyle(t *testing.T) {
	s1 := newMockShape(model.Frame{}, nil, style.NewStyle(true, model.Red))
	s2 := newMockShape(model.Frame{}, nil, style.NewStyle(true, model.Blue))

	g := NewGroup()
	g.AddShape(s1)
	g.AddShape(s2)

	assert.Equal(t, model.Undefined, g.GetFillStyle().GetColor())
	assert.True(t, g.GetFillStyle().IsEnabled())
}

func TestGroup_SetGroupStyle(t *testing.T) {
	s1 := newMockShape(model.Frame{}, nil, style.NewStyle(true, model.Red))
	s2 := newMockShape(model.Frame{}, nil, style.NewStyle(true, model.Blue))

	g := NewGroup()
	g.AddShape(s1)
	g.AddShape(s2)

	newStyle := style.NewStyle(true, model.Green)
	g.SetFillStyle(newStyle)

	assert.Equal(t, model.Green, s1.GetFillStyle().GetColor())
	assert.Equal(t, model.Green, s2.GetFillStyle().GetColor())
}
