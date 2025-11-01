package command

import "editor/pkg/document"

type setTitleCommand struct {
	AbstractCommand
	doc      document.Document
	newTitle string
	oldTitle string
}

func NewSetTitleCommand(doc document.Document, newTitle string) Command {
	return &setTitleCommand{
		doc:      doc,
		newTitle: newTitle,
	}
}

func (c *setTitleCommand) Execute() {
	c.oldTitle = c.doc.GetTitle()
	c.doc.SetTitle(c.newTitle)
}

func (c *setTitleCommand) Unexecute() {
	c.doc.SetTitle(c.oldTitle)
}

func (c *setTitleCommand) Merge(next Command) bool {
	nextCmd, ok := next.(*setTitleCommand)
	if ok {
		c.newTitle = nextCmd.newTitle
		return true
	}
	return false
}
