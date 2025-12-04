package command

import "gumballmachine/pkg/gumballmachine"

type TurnCrankCommand struct {
	machine *gumballmachine.GumballMachine
}

func NewTurnCrankCommand(m *gumballmachine.GumballMachine) *TurnCrankCommand {
	return &TurnCrankCommand{machine: m}
}

func (c *TurnCrankCommand) Execute() {
	c.machine.TurnCrank()
}
