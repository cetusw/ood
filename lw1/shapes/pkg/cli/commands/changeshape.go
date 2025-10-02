package commands

import (
	"fmt"
	"shapes/pkg/shapes"

	"shapes/pkg/cli/parser"
)

func ChangeShapeCommand(picture *shapes.Picture, args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("недостаточно аргументов для ChangeShape. Ожидалось: <shapeID>, <тип>, <параметры>")
	}

	shapeID := args[0]
	shapeString := args[1]

	strategy, err := parser.StrategyInterpreter(shapeString, args[2:])
	if err != nil {
		return err
	}
	picture.ChangeShape(shapeID, strategy)

	return nil
}
