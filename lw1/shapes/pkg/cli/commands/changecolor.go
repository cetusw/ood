package commands

import (
	"fmt"
	"shapes/pkg/shapes"
)

func ChangeColorCommand(picture *shapes.Picture, args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("недостаточно аргументов для ChangeColor. Ожидалось: <id> <цвет>")
	}

	id := args[0]
	color := args[1]

	picture.ChangeColor(id, color)

	return nil
}
