package factory

import (
	"factory/pkg/domain"
	"factory/pkg/shapes"
	"fmt"
	"strconv"
	"strings"
)

type ShapeFactory interface {
	CreateShape(description string) (domain.Shape, error)
}

type shapeFactory struct{}

func NewShapeFactory() ShapeFactory {
	return &shapeFactory{}
}

func (f *shapeFactory) CreateShape(description string) (domain.Shape, error) {
	parts := strings.Fields(description)
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid description: %s", description)
	}

	shapeType := parts[0]
	color := domain.Color(parts[1])
	args := parts[2:]

	switch shapeType {
	case "rectangle":
		nums, err := toInts(args, 4)
		if err != nil {
			return nil, err
		}
		p1 := domain.Point{X: nums[0], Y: nums[1]}
		p2 := domain.Point{X: nums[2], Y: nums[3]}
		return shapes.NewRectangle(color, p1, p2), nil
	case "triangle":
		nums, err := toInts(args, 6)
		if err != nil {
			return nil, err
		}
		v1 := domain.Point{X: nums[0], Y: nums[1]}
		v2 := domain.Point{X: nums[2], Y: nums[3]}
		v3 := domain.Point{X: nums[4], Y: nums[5]}
		return shapes.NewTriangle(color, v1, v2, v3), nil
	case "ellipse":
		nums, err := toInts(args, 4)
		if err != nil {
			return nil, err
		}
		center := domain.Point{X: nums[0], Y: nums[1]}
		return shapes.NewEllipse(color, center, nums[2], nums[3]), nil
	case "polygon":
		nums, err := toInts(args, 4)
		if err != nil {
			return nil, err
		}
		center := domain.Point{X: nums[0], Y: nums[1]}
		return shapes.NewPolygon(color, center, nums[2], nums[3]), nil
	default:
		return nil, fmt.Errorf("unknown shape type: %s", shapeType)
	}
}

func toInts(s []string, expectedLen int) ([]int, error) {
	if len(s) != expectedLen {
		return nil, fmt.Errorf("expected %d arguments, got %d", expectedLen, len(s))
	}
	nums := make([]int, len(s))
	for i, v := range s {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("failed to parse number: %s", v)
		}
		nums[i] = num
	}
	return nums, nil
}
