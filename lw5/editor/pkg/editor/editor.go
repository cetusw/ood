package editor

import (
	"bufio"
	"editor/pkg/document"
	"editor/pkg/factory"
	"editor/pkg/history"
	"fmt"
	"os"
	"strings"
)

type Editor interface {
	Run()
}

type editor struct {
	doc             document.Document
	documentHistory history.History
	commandFactory  factory.CommandFactory
	reader          *bufio.Scanner
}

func NewEditor() Editor {
	doc := document.NewDocument("Untitled Document")
	return &editor{
		doc:             doc,
		documentHistory: history.NewHistory(),
		commandFactory:  factory.NewCommandFactory(),
		reader:          bufio.NewScanner(os.Stdin),
	}
}

func (e *editor) Run() {
	fmt.Println("Type 'Help' to get list of commands.")
	for {
		fmt.Print("> ")
		if !e.reader.Scan() {
			break
		}

		line := e.reader.Text()
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		commandName := parts[0]

		switch commandName {
		case "Help":
			printHelp()
			continue
		case "List":
			e.doc.List()
			continue
		case "Undo":
			e.documentHistory.Undo()
			continue
		case "Redo":
			e.documentHistory.Redo()
			continue
		case "Save":
			if len(parts) < 2 {
				fmt.Println("usage: Save <path>")
				continue
			}
			if err := e.doc.Save(parts[1]); err != nil {
				fmt.Printf("Failed to save: %v\n", err)
			} else {
				fmt.Println("Saved successfully")
			}
			continue
		case "Exit":
			fmt.Println("Bye")
			return
		}

		cmd, err := e.commandFactory.CreateCommand(line, e.doc)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		if cmd != nil {
			e.documentHistory.AddAndExecuteCommand(cmd)
		}
	}
}

func printHelp() {
	fmt.Println("\nList of available commands:")
	fmt.Println("  SetTitle <title>                                      - Set document title.")
	fmt.Println("  InsertParagraph <position|end> <text>                 - Insert paragraph.")
	fmt.Println("  InsertImage <position|end> <width> <height> <path>    - Insert image.")
	fmt.Println("  ReplaceText <position> <text>                         - Replace text in paragraph.")
	fmt.Println("  ResizeImage <position> <new_width> <new_height>       - Resize image.")
	fmt.Println("  DeleteItem <position>                                 - Delete item.")
	fmt.Println("  List                                                  - List items.")
	fmt.Println("  Undo                                                  - Undo last command.")
	fmt.Println("  Redo                                                  - Redo last command.")
	fmt.Println("  Save <filepath.html>                                  - Save document to HTML.")
	fmt.Println("  Exit                                                  - Exit.")
	fmt.Println()
}
