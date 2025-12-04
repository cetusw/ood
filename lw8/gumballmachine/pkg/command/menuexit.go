package command

type ExitMenuCommand struct {
	menu Menu
}

func NewExitMenuCommand(m Menu) *ExitMenuCommand {
	return &ExitMenuCommand{menu: m}
}

func (c *ExitMenuCommand) Execute() {
	c.menu.Exit()
}
