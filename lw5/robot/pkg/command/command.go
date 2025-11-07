package command

type Command interface {
	Execute()
}

type Menu interface {
	ShowInstructions()
	Exit()
}
