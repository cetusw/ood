package main

import (
	"robot/pkg/command"
	"robot/pkg/menu"
	"robot/pkg/robot"
)

func main() {
	r := robot.NewRobot()
	m := menu.NewMenu()

	m.AddItem("on", "Turns the Robot on", command.NewTurnOnCommand(r))
	m.AddItem("off", "Turns the Robot off", command.NewTurnOffCommand(r))
	m.AddItem("north", "Makes the Robot walk north", command.NewWalkCommand(r, robot.North))
	m.AddItem("south", "Makes the Robot walk south", command.NewWalkCommand(r, robot.South))
	m.AddItem("west", "Makes the Robot walk west", command.NewWalkCommand(r, robot.West))
	m.AddItem("east", "Makes the Robot walk east", command.NewWalkCommand(r, robot.East))
	m.AddItem("stop", "Stops the Robot", command.NewStopCommand(r))

	macro := command.NewMacroCommand()
	macro.AddCommand(command.NewTurnOnCommand(r))
	macro.AddCommand(command.NewWalkCommand(r, robot.North))
	macro.AddCommand(command.NewWalkCommand(r, robot.East))
	macro.AddCommand(command.NewWalkCommand(r, robot.South))
	macro.AddCommand(command.NewWalkCommand(r, robot.West))
	macro.AddCommand(command.NewTurnOffCommand(r))
	m.AddItem("patrol", "Patrol the territory", macro)

	m.AddItem("help", "Show instructions", command.NewMenuHelpCommand(m))
	m.AddItem("exit", "Exit from this menu", command.NewExitMenuCommand(m))

	m.Run()
}
