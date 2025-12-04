package command

import (
	"fmt"
	"gumballmachine/pkg/gumballmachine"
)

type InfoCommand struct {
	machine *gumballmachine.GumballMachine
}

func NewInfoCommand(m *gumballmachine.GumballMachine) *InfoCommand {
	return &InfoCommand{machine: m}
}

func (c *InfoCommand) Execute() {
	fmt.Println(c.machine.String())
}
