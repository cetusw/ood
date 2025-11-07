package command

import "robot/pkg/robot"

type TurnOffCommand struct {
	robot *robot.Robot
}

func NewTurnOffCommand(r *robot.Robot) *TurnOffCommand {
	return &TurnOffCommand{robot: r}
}

func (c *TurnOffCommand) Execute() {
	c.robot.TurnOff()
}
