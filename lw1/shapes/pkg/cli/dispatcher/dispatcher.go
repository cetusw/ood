package dispatcher

import (
	"fmt"
	"shapes/pkg/shapes"
	"strings"
)

type CommandFunc func(picture *shapes.Picture, args []string) error

type CommandDispatcher struct {
	commands map[string]CommandFunc
}

func NewDispatcher() *CommandDispatcher {
	return &CommandDispatcher{
		commands: make(map[string]CommandFunc),
	}
}

func (d *CommandDispatcher) Register(name string, fn CommandFunc) {
	d.commands[strings.ToLower(name)] = fn
}

func (d *CommandDispatcher) Execute(picture *shapes.Picture, commandName string, args []string) error {
	cmd, exists := d.commands[strings.ToLower(commandName)]
	if !exists {
		return fmt.Errorf("unknown command: %s", commandName)
	}
	return cmd(picture, args)
}
