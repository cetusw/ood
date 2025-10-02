package commands

import (
	"fmt"
	"shapes/pkg/shapes"

	"shapes/pkg/cli/parser"
	"shapes/pkg/shapes/shape"
)

func AddShapeCommand(picture *shapes.Picture, args []string) error {
	if len(args) < 4 {
		return fmt.Errorf("недостаточно аргументов для AddShape. Ожидалось: <shapeID>, <color>, <type>, <params>")
	}

	shapeID := args[0]
	color := args[1]
	shapeString := args[2]

	strategy, err := parser.StrategyInterpreter(shapeString, args[3:])
	if err != nil {
		return err
	}
	s := shape.NewShape(strategy, shapeID, color)
	err = picture.AddShape(s)
	if err != nil {
		return err
	}

	return nil
}
