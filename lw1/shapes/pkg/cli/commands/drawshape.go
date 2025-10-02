package commands

import (
	"fmt"
	"shapes/pkg/common"
	"shapes/pkg/shapes"
)

func DrawShapeCommand(picture *shapes.Picture, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("недостаточно аргументов для DrawShape. Ожидалось: <shapeID>")
	}

	shapeID := args[0]

	picture.DrawShape(shapeID, picture.GetCanvas())
	err := picture.GetCanvas().SaveToFile(common.PictureFileName)
	if err != nil {
		fmt.Println("Ошибка сохранения в файл:", err)
	}

	return nil
}
