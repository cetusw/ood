package commands

import (
	"fmt"
	"shapes/pkg/shapes"
	"shapes/pkg/shapes/model"
	"strconv"
)

func MovePictureCommand(picture *shapes.Picture, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("недостаточно аргументов для MovePicture. Ожидалось: <dx> <dy>")
	}

	dxString := args[0]
	dyString := args[1]

	dx, err := strconv.ParseFloat(dxString, 64)
	if err != nil {
		return fmt.Errorf("неверный формат для смещения dx: %v", err)
	}
	dy, err := strconv.ParseFloat(dyString, 64)
	if err != nil {
		return fmt.Errorf("неверный формат для смещения dy: %v", err)
	}

	vector := model.Point{X: dx, Y: dy}
	picture.MovePicture(vector)

	return nil
}
