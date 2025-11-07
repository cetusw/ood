package command

import "robot/pkg/robot"

type TurnOnCommand struct {
	robot *robot.Robot
}

func NewTurnOnCommand(r *robot.Robot) *TurnOnCommand {
	return &TurnOnCommand{robot: r}
}

func (c *TurnOnCommand) Execute() {
	c.robot.TurnOn()
}
