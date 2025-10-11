package shapefactory

import (
	"fmt"
	"strconv"
	"strings"

	"factory/pkg/model"
	"factory/pkg/shapes"
)

type ShapeFactory interface {
	CreateShape(description string) (shapes.Shape, error)
}

type shapeFactory struct{}

func NewShapeFactory() ShapeFactory {
	return &shapeFactory{}
}

func (f *shapeFactory) CreateShape(description string) (shapes.Shape, error) {
	parts := strings.Fields(description)
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid description: %s", description)
	}

	shapeType := parts[0]
	color := model.Color(parts[1])
	args := parts[2:]

	switch shapeType {
	case "rectangle":
		nums, err := toFloats(args, 4)
		if err != nil {
			return nil, err
		}
		p1 := model.Point{X: nums[0], Y: nums[1]}
		p2 := model.Point{X: nums[2], Y: nums[3]}
		return shapes.NewRectangle(color, p1, p2), nil
	case "triangle":
		nums, err := toFloats(args, 6)
		if err != nil {
			return nil, err
		}
		v1 := model.Point{X: nums[0], Y: nums[1]}
		v2 := model.Point{X: nums[2], Y: nums[3]}
		v3 := model.Point{X: nums[4], Y: nums[5]}
		return shapes.NewTriangle(color, v1, v2, v3), nil
	case "ellipse":
		nums, err := toFloats(args, 4)
		if err != nil {
			return nil, err
		}
		center := model.Point{X: nums[0], Y: nums[1]}
		radius := model.Radius{X: nums[2], Y: nums[3]}
		return shapes.NewEllipse(color, center, radius), nil
	case "polygon":
		nums, err := toFloats(args, 4)
		if err != nil {
			return nil, err
		}
		center := model.Point{X: nums[0], Y: nums[1]}
		return shapes.NewPolygon(color, center, nums[2], int(nums[3])), nil
	default:
		return nil, fmt.Errorf("unknown shape type: %s", shapeType)
	}
}

func toFloats(str []string, expectedLen int) ([]float64, error) {
	if len(str) != expectedLen {
		return nil, fmt.Errorf("expected %d arguments, got %d", expectedLen, len(str))
	}
	nums := make([]float64, len(str))
	for i, value := range str {
		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, fmt.Errorf("неверный формат для координаты X: %v", err)
		}
		nums[i] = num
	}
	return nums, nil
}
