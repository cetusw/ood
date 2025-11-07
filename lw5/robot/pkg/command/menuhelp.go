package command

type MenuHelpCommand struct {
	menu Menu
}

func NewMenuHelpCommand(m Menu) *MenuHelpCommand {
	return &MenuHelpCommand{menu: m}
}

func (c *MenuHelpCommand) Execute() {
	c.menu.ShowInstructions()
}
