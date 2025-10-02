package commands

import (
	"fmt"
	"shapes/pkg/shapes"
	"shapes/pkg/shapes/model"
	"strconv"
)

func MoveShapeCommand(picture *shapes.Picture, args []string) error {
	if len(args) < 3 {
		return fmt.Errorf("недостаточно аргументов для MoveShape. Ожидалось: <shapeID> <dx> <dy>")
	}

	shapeID := args[0]
	dxString := args[1]
	dyString := args[2]

	dx, err := strconv.ParseFloat(dxString, 64)
	if err != nil {
		return fmt.Errorf("неверный формат для смещения dx: %v", err)
	}
	dy, err := strconv.ParseFloat(dyString, 64)
	if err != nil {
		return fmt.Errorf("неверный формат для смещения dy: %v", err)
	}

	vector := model.Point{X: dx, Y: dy}
	picture.MoveShape(shapeID, vector)

	return nil
}
