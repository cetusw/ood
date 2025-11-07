package command

import "robot/pkg/robot"

type StopCommand struct {
	robot *robot.Robot
}

func NewStopCommand(r *robot.Robot) *StopCommand {
	return &StopCommand{robot: r}
}

func (c *StopCommand) Execute() {
	c.robot.Stop()
}
