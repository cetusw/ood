package command

import (
	"fmt"
	"gumballmachine/pkg/gumballmachine"
)

type RefillCommand struct {
	machine *gumballmachine.GumballMachine
}

func NewRefillCommand(m *gumballmachine.GumballMachine) *RefillCommand {
	return &RefillCommand{machine: m}
}

func (c *RefillCommand) Execute() {
	var count uint
	fmt.Print("Enter amount to refill: ")
	_, err := fmt.Scanln(&count)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}
	c.machine.Refill(count)
	fmt.Printf("Machine refilled. Current inventory: %d\n", c.machine.GetBallCount())
}
