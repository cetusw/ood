package history

import (
	"editor/pkg/command"
	"fmt"
)

const maxHistorySize = 10

type History interface {
	AddAndExecuteCommand(cmd command.Command)
	CanUndo() bool
	Undo()
	CanRedo() bool
	Redo()
}

type history struct {
	commands         []command.Command
	nextCommandIndex int
}

func NewHistory() History {
	return &history{
		commands:         make([]command.Command, 0, maxHistorySize),
		nextCommandIndex: 0,
	}
}

func (h *history) AddAndExecuteCommand(cmd command.Command) {
	if h.nextCommandIndex < len(h.commands) {
		for _, discardedCmd := range h.commands[h.nextCommandIndex:] {
			discardedCmd.Destroy()
		}
		h.commands = h.commands[:h.nextCommandIndex]
	}

	cmd.Execute()

	if h.nextCommandIndex > 0 {
		lastCmd := h.commands[h.nextCommandIndex-1]
		if lastCmd.Merge(cmd) {
			return
		}
	}

	if len(h.commands) == maxHistorySize {
		h.commands = h.commands[1:]
		h.nextCommandIndex--
	}

	h.commands = append(h.commands, cmd)
	h.nextCommandIndex++
}

func (h *history) CanUndo() bool {
	return h.nextCommandIndex > 0
}

func (h *history) Undo() {
	if !h.CanUndo() {
		fmt.Println("Cannot undo")
		return
	}
	h.nextCommandIndex--
	cmd := h.commands[h.nextCommandIndex]
	cmd.Unexecute()
	fmt.Println("Undo successfully")
}

func (h *history) CanRedo() bool {
	return h.nextCommandIndex < len(h.commands)
}

func (h *history) Redo() {
	if !h.CanRedo() {
		fmt.Println("Cannot redo")
		return
	}
	cmd := h.commands[h.nextCommandIndex]
	cmd.Execute()
	h.nextCommandIndex++
	fmt.Println("Redo successfully")
}
