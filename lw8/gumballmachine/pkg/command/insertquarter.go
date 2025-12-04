package command

import "gumballmachine/pkg/gumballmachine"

type InsertQuarterCommand struct {
	machine *gumballmachine.GumballMachine
}

func NewInsertQuarterCommand(m *gumballmachine.GumballMachine) *InsertQuarterCommand {
	return &InsertQuarterCommand{machine: m}
}

func (c *InsertQuarterCommand) Execute() {
	c.machine.InsertQuarter()
}
