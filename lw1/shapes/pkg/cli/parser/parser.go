package parser

import (
	"fmt"
	"shapes/pkg/common"
	"shapes/pkg/shapes/model"
	"shapes/pkg/shapes/shape"
	"strconv"
	"strings"
)

type Command struct {
	Name      string
	Arguments []string
}

func ParseCommand(input string) (Command, error) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return Command{}, fmt.Errorf("command cannot be empty")
	}

	cmdName := parts[0]
	args := parts[1:]

	return Command{
		Name:      cmdName,
		Arguments: args,
	}, nil
}

func StrategyInterpreter(shapeString string, args []string) (shape.Strategy, error) {
	switch shapeString {
	case common.Circle:
		strategy, err := parseCircleArgs(args)
		if err != nil {
			return &shape.CircleStrategy{}, err
		}
		return strategy, nil
	case common.Rectangle:
		strategy, err := parseRectangleArgs(args)
		if err != nil {
			return &shape.RectangleStrategy{}, err
		}
		return strategy, nil
	case common.Triangle:
		strategy, err := parseTriangleArgs(args)
		if err != nil {
			return &shape.TriangleStrategy{}, err
		}
		return strategy, nil
	case common.Line:
		strategy, err := parseLineArgs(args)
		if err != nil {
			return &shape.LineStrategy{}, err
		}
		return strategy, nil
	case common.Text:
		strategy, err := parseTextArgs(args)
		if err != nil {
			return &shape.TextStrategy{}, err
		}
		return strategy, nil
	default:
		return nil, fmt.Errorf("unknown shape type: %s", shapeString)
	}
}

func parseCircleArgs(args []string) (*shape.CircleStrategy, error) {
	if len(args) < 3 {
		return &shape.CircleStrategy{}, fmt.Errorf("недостаточно аргументов для отрисовки %s", common.Circle)
	}
	centerX, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return nil, fmt.Errorf("неверный формат для координаты X: %v", err)
	}
	centerY, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return nil, fmt.Errorf("неверный формат для координаты Y: %v", err)
	}
	radius, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return nil, fmt.Errorf("неверный формат для радиуса: %v", err)
	}

	return shape.NewCircleStrategy(model.Point{
		X: centerX,
		Y: centerY,
	}, radius), nil
}

func parseRectangleArgs(args []string) (*shape.RectangleStrategy, error) {
	if len(args) < 4 {
		return &shape.RectangleStrategy{}, fmt.Errorf("недостаточно аргументов для отрисовки %s", common.Circle)
	}
	left, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return nil, fmt.Errorf("неверный формат для координаты X: %v", err)
	}
	top, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return nil, fmt.Errorf("неверный формат для координаты Y: %v", err)
	}
	width, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return nil, fmt.Errorf("неверный формат для длины: %v", err)
	}
	height, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return nil, fmt.Errorf("неверный формат для высоты: %v", err)
	}

	return shape.NewRectangleStrategy(model.Point{
		X: left,
		Y: top,
	}, width, height), nil
}

func parseTriangleArgs(args []string) (*shape.TriangleStrategy, error) {
	if len(args) < 6 {
		return &shape.TriangleStrategy{}, fmt.Errorf("недостаточно аргументов для отрисовки %s", common.Rectangle)
	}
	var err error
	var vertices [3]model.Point
	for i := 0; i < len(vertices); i++ {
		vertices[i].X, err = strconv.ParseFloat(args[i*2], 64)
		if err != nil {
			return nil, fmt.Errorf("неверный формат для координаты X: %v", err)
		}
		vertices[i].Y, err = strconv.ParseFloat(args[(i*2)+1], 64)
		if err != nil {
			return nil, fmt.Errorf("неверный формат для координаты Y: %v", err)
		}
	}

	return shape.NewTriangleStrategy(vertices), nil
}

func parseLineArgs(args []string) (*shape.LineStrategy, error) {
	if len(args) < 4 {
		return &shape.LineStrategy{}, fmt.Errorf("недостаточно аргументов для отрисовки %s", common.Line)
	}
	var err error
	var vertices [2]model.Point
	for i := 0; i < len(vertices); i++ {
		vertices[i].X, err = strconv.ParseFloat(args[i*2], 64)
		if err != nil {
			return nil, fmt.Errorf("неверный формат для координаты X: %v", err)
		}
		vertices[i].Y, err = strconv.ParseFloat(args[(i*2)+1], 64)
		if err != nil {
			return nil, fmt.Errorf("неверный формат для координаты Y: %v", err)
		}
	}

	return shape.NewLineStrategy(vertices), nil
}

func parseTextArgs(args []string) (*shape.TextStrategy, error) {
	if len(args) < 4 {
		return &shape.TextStrategy{}, fmt.Errorf("недостаточно аргументов для отрисовки %s", common.Text)
	}
	left, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return nil, fmt.Errorf("неверный формат для координаты X: %v", err)
	}
	top, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return nil, fmt.Errorf("неверный формат для координаты Y: %v", err)
	}
	size, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return nil, fmt.Errorf("неверный формат для длины: %v", err)
	}
	content := args[3]

	return shape.NewTextStrategy(model.Point{
		X: left,
		Y: top,
	}, size, content), nil
}
