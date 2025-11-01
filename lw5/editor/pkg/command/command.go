package command

type Command interface {
	Execute()
	Unexecute()
	Destroy()
	Merge(next Command) bool
}

type AbstractCommand struct{}

func (c *AbstractCommand) Destroy() {}

func (c *AbstractCommand) Execute() {}

func (c *AbstractCommand) Unexecute() {}

func (c *AbstractCommand) Merge(_ Command) bool {
	return false
}
