package commands

import (
	"shapes/pkg/shapes"
)

func ListCommand(picture *shapes.Picture, args []string) error {

	picture.ListShapes()

	return nil
}
