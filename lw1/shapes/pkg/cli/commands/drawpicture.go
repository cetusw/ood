package commands

import (
	"fmt"
	"shapes/pkg/common"
	"shapes/pkg/shapes"
)

func DrawPictureCommand(picture *shapes.Picture, args []string) error {
	picture.DrawPicture(picture.GetCanvas())
	err := picture.GetCanvas().SaveToFile(common.PictureFileName)
	if err != nil {
		fmt.Println("Ошибка сохранения в файл:", err)
	}

	return nil
}
