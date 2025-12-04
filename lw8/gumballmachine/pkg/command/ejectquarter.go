package command

import "gumballmachine/pkg/gumballmachine"

type EjectQuarterCommand struct {
	machine *gumballmachine.GumballMachine
}

func NewEjectQuarterCommand(m *gumballmachine.GumballMachine) *EjectQuarterCommand {
	return &EjectQuarterCommand{machine: m}
}

func (c *EjectQuarterCommand) Execute() {
	c.machine.EjectQuarter()
}
