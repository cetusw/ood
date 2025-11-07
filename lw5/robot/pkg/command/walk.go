package command

import "robot/pkg/robot"

type WalkCommand struct {
	robot     *robot.Robot
	direction robot.WalkDirection
}

func NewWalkCommand(r *robot.Robot, dir robot.WalkDirection) *WalkCommand {
	return &WalkCommand{robot: r, direction: dir}
}

func (c *WalkCommand) Execute() {
	c.robot.Walk(c.direction)
}
