package common

import (
	"fmt"
	"image/color"
	"shapes/pkg/shapes/model"
)

func ParseHexColor(hex string) (color.RGBA, error) {
	var r, g, b uint8
	if len(hex) == 7 && hex[0] == '#' {
		_, err := fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
		if err == nil {
			return color.RGBA{R: r, G: g, B: b, A: 255}, nil
		}
		return color.RGBA{}, err
	}
	return color.RGBA{}, fmt.Errorf("некорректный hex: %s", hex)
}

func MovePoint(point model.Point, vector model.Point) model.Point {
	return model.Point{
		X: point.X + vector.X,
		Y: point.Y + vector.Y,
	}
}
