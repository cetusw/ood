package main

import (
	"bufio"
	"fmt"
	"os"
	"shapes/pkg/cli/commands"
	"shapes/pkg/cli/dispatcher"
	"shapes/pkg/cli/parser"
	"shapes/pkg/shapes"
)

func main() {
	d := dispatcher.NewDispatcher()
	registerCommands(d)
	picture := shapes.NewPicture()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите команду (например, AddShape circ #febb38 circle 100 200 25):")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" {
			break
		}

		cmd, err := parser.ParseCommand(input)
		if err != nil {
			fmt.Println("Ошибка парсинга:", err)
			continue
		}
		err = d.Execute(picture, cmd.Name, cmd.Arguments)
		if err != nil {
			fmt.Println("Ошибка выполнения команды:", err)
		}

		fmt.Println("\nВведите следующую команду:")
	}
}

func registerCommands(d *dispatcher.CommandDispatcher) {
	d.Register("AddShape", commands.AddShapeCommand)
	d.Register("MoveShape", commands.MoveShapeCommand)
	d.Register("MovePicture", commands.MovePictureCommand)
	d.Register("DeleteShape", commands.DeleteShapeCommand)
	d.Register("List", commands.ListCommand)
	d.Register("ChangeColor", commands.ChangeColorCommand)
}
