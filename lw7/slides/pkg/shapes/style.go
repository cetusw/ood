package shapes

import "slides/pkg/model"

type Style interface {
	Clone() Style
	IsEnabled() bool
	Enable(enable bool)
	GetColor() model.Color
	SetColor(color model.Color)
}

type style struct {
	isEnabled bool
	color     model.Color
}

func NewStyle(isEnabled bool, color model.Color) Style {
	return &style{
		isEnabled: isEnabled,
		color:     color,
	}
}

func (s *style) Clone() Style {
	return &style{
		isEnabled: s.isEnabled,
		color:     s.color,
	}
}

func (s *style) IsEnabled() bool {
	return s.isEnabled
}

func (s *style) Enable(enable bool) {
	s.isEnabled = enable
}

func (s *style) GetColor() model.Color {
	return s.color
}

func (s *style) SetColor(color model.Color) {
	s.color = color
}
