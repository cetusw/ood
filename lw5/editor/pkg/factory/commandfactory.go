package factory

import (
	"editor/pkg/command"
	"editor/pkg/document"
	"editor/pkg/model"
	"fmt"
	"strconv"
	"strings"
)

const (
	end = "end"
)

type CommandFactory interface {
	CreateCommand(input string, doc document.Document) (command.Command, error)
}

type commandFactory struct{}

func NewCommandFactory() CommandFactory {
	return &commandFactory{}
}

func (f *commandFactory) CreateCommand(input string, doc document.Document) (command.Command, error) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return nil, fmt.Errorf("command is empty")
	}

	commandName := parts[0]
	args := parts[1:]

	switch commandName {
	case "InsertParagraph":
		if len(args) < 2 {
			return nil, fmt.Errorf("usage: InsertParagraph <position|end> <text>")
		}
		pos, err := parsePosition(args[0], doc.GetItemsCount(), true)
		if err != nil {
			return nil, err
		}
		text := args[1]
		return command.NewInsertParagraphCommand(doc, pos, text), nil

	case "InsertImage":
		if len(args) != 4 {
			return nil, fmt.Errorf("usage: InsertImage <position|end> <width> <height> <path>")
		}
		pos, err := parsePosition(args[0], doc.GetItemsCount(), true)
		if err != nil {
			return nil, err
		}
		width, err := toInt(args[1])
		if err != nil {
			return nil, err
		}
		height, err := toInt(args[2])
		if err != nil {
			return nil, err
		}
		return command.NewInsertImageCommand(doc, args[3], model.Size{Width: width, Height: height}, pos), nil

	case "SetTitle":
		if len(args) == 0 {
			return nil, fmt.Errorf("usage: SetTitle <new_title>")
		}
		title := args[0]
		return command.NewSetTitleCommand(doc, title), nil

	case "ReplaceText":
		if len(args) < 2 {
			return nil, fmt.Errorf("usage: ReplaceText <position> <new_text>")
		}
		pos, err := parsePosition(args[0], doc.GetItemsCount(), false)
		if err != nil {
			return nil, err
		}
		text := args[1]
		return command.NewReplaceTextCommand(doc, pos, text), nil

	case "ResizeImage":
		if len(args) != 3 {
			return nil, fmt.Errorf("usage: ResizeImage <position> <width> <height>")
		}
		pos, err := parsePosition(args[0], doc.GetItemsCount(), false)
		if err != nil {
			return nil, err
		}
		width, err := toInt(args[1])
		if err != nil {
			return nil, err
		}
		height, err := toInt(args[2])
		if err != nil {
			return nil, err
		}
		return command.NewResizeImageCommand(doc, pos, model.Size{
			Width:  width,
			Height: height,
		}), nil

	case "DeleteItem":
		if len(args) != 1 {
			return nil, fmt.Errorf("usage: DeleteItem <position>")
		}
		pos, err := parsePosition(args[0], doc.GetItemsCount(), false)
		if err != nil {
			return nil, err
		}
		return command.NewDeleteItemCommand(doc, pos), nil

	case "List", "Help", "Undo", "Redo", "Save", "Exit":
		return nil, nil

	default:
		return nil, fmt.Errorf("unknown command: %s", commandName)
	}
}

func parsePosition(value string, count int, allowEnd bool) (int, error) {
	if value == end {
		if allowEnd {
			return count, nil
		}
		return 0, fmt.Errorf("position 'end' restricted for this command")
	}

	num, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid position %s", value)
	}
	if num < 1 || num > count {
		return 0, fmt.Errorf("position %d out of range [1..%d]", num, count)
	}
	return num - 1, nil
}

func toInt(value string) (int, error) {
	num, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid integer '%s'", value)
	}
	return num, nil
}
