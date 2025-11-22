package shapes

import (
	"math"
	"slides/pkg/canvas"
	"slides/pkg/model"
)

type Group struct {
	shapes []Shape
}

func NewGroup() *Group {
	return &Group{shapes: make([]Shape, 0)}
}

func (g *Group) Clone() Shape {
	newGroup := NewGroup()

	for _, s := range g.shapes {
		newGroup.AddShape(s.Clone())
	}

	return newGroup
}

func (g *Group) AddShape(s Shape) {
	g.shapes = append(g.shapes, s)
}

func (g *Group) Draw(c canvas.Canvas) {
	for _, s := range g.shapes {
		s.Draw(c)
	}
}

func (g *Group) GetFrame() model.Frame {
	if len(g.shapes) == 0 {
		return model.Frame{}
	}
	leftX, topY := math.MaxFloat64, -math.MaxFloat64
	rightX, bottomY := -math.MaxFloat64, math.MaxFloat64

	for _, s := range g.shapes {
		f := s.GetFrame()
		leftX = math.Min(leftX, f.X)
		topY = math.Max(topY, f.Y)
		rightX = math.Max(rightX, f.X+f.Width)
		bottomY = math.Min(bottomY, f.Y-f.Height)
	}
	return model.Frame{
		Point: model.Point{
			X: leftX,
			Y: topY,
		},
		Width:  rightX - leftX,
		Height: topY - bottomY,
	}
}

func (g *Group) SetFrame(frame model.Frame) {
	groupFrame := g.GetFrame()
	if groupFrame.Width == 0 || groupFrame.Height == 0 {
		return
	}
	scaleX := frame.Width / groupFrame.Width
	scaleY := frame.Height / groupFrame.Height

	for _, s := range g.shapes {
		shapeFrame := s.GetFrame()
		newLeft := frame.X + (shapeFrame.X-groupFrame.X)*scaleX
		newTop := frame.Y + (shapeFrame.Y-groupFrame.Y)*scaleY
		newWidth := shapeFrame.Width * scaleX
		newHeight := shapeFrame.Height * scaleY
		s.SetFrame(model.Frame{
			Point: model.Point{
				X: newLeft,
				Y: newTop,
			},
			Width:  newWidth,
			Height: newHeight,
		})
	}
}

func (g *Group) GetFillStyle() Style {
	if len(g.shapes) == 0 {
		return NewStyle(false, model.Undefined)
	}

	firstStyle := g.shapes[0].GetFillStyle()
	commonColor := firstStyle.GetColor()
	commonEnabled := firstStyle.IsEnabled()

	for i := 1; i < len(g.shapes); i++ {
		currStyle := g.shapes[i].GetFillStyle()

		if currStyle.GetColor() != commonColor {
			commonColor = model.Undefined
		}
		if currStyle.IsEnabled() != commonEnabled {
			commonEnabled = false
		}
		if commonColor == model.Undefined && commonEnabled == false {
			break
		}
	}

	return NewStyle(commonEnabled, commonColor)
}

func (g *Group) SetFillStyle(fillStyle Style) {
	for _, s := range g.shapes {
		currentStyle := s.GetFillStyle()

		if fillStyle.GetColor() != model.Undefined {
			currentStyle.SetColor(fillStyle.GetColor())
		}

		currentStyle.Enable(fillStyle.IsEnabled())
		s.SetFillStyle(currentStyle)
	}
}

func (g *Group) GetLineStyle() Style {
	if len(g.shapes) == 0 {
		return NewStyle(false, model.Undefined)
	}

	firstStyle := g.shapes[0].GetLineStyle()
	commonColor := firstStyle.GetColor()
	commonEnabled := firstStyle.IsEnabled()

	for i := 1; i < len(g.shapes); i++ {
		currStyle := g.shapes[i].GetLineStyle()

		if currStyle.GetColor() != commonColor {
			commonColor = model.Undefined
		}
		if currStyle.IsEnabled() != commonEnabled {
			commonEnabled = false
		}
		if commonColor == model.Undefined && commonEnabled == false {
			break
		}
	}

	return NewStyle(commonEnabled, commonColor)
}

func (g *Group) SetLineStyle(lineStyle Style) {
	for _, s := range g.shapes {
		currentStyle := s.GetLineStyle()

		if lineStyle.GetColor() != model.Undefined {
			currentStyle.SetColor(lineStyle.GetColor())
		}
		currentStyle.Enable(lineStyle.IsEnabled())

		s.SetLineStyle(currentStyle)
	}
}
