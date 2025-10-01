package commands

import (
	"fmt"
	"shapes/pkg/shapes"
)

func DeleteShapeCommand(picture *shapes.Picture, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("недостаточно аргументов для DeleteShape. Ожидалось: <id>")
	}

	id := args[0]

	picture.DeleteShape(id)

	return nil
}
