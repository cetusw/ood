package main

import (
	"gumballmachine/pkg/command"
	"gumballmachine/pkg/gumballmachine"
	"gumballmachine/pkg/menu"
	"os"
)

const defaultBallsCount = 10

func main() {
	machine := gumballmachine.NewGumballMachine(defaultBallsCount, os.Stdout)
	m := menu.NewMenu()

	m.AddItem("insert", "Inserts coin into machine", command.NewInsertQuarterCommand(machine))
	m.AddItem("turn", "Turns the crank", command.NewTurnCrankCommand(machine))
	m.AddItem("eject", "Ejects coins", command.NewEjectQuarterCommand(machine))
	m.AddItem("refill", "Refills balls", command.NewRefillCommand(machine))
	m.AddItem("info", "Shows machine information", command.NewInfoCommand(machine))

	m.AddItem("help", "Show instructions", command.NewMenuHelpCommand(m))
	m.AddItem("exit", "Exit from this menu", command.NewExitMenuCommand(m))

	m.Run()
}
