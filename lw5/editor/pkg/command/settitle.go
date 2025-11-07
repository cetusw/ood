package command

import "editor/pkg/document"

type setTitleCommand struct {
	AbstractCommand
	doc      document.Document
	NewTitle string
	OldTitle string
}

func NewSetTitleCommand(doc document.Document, newTitle string) Command {
	return &setTitleCommand{
		doc:      doc,
		NewTitle: newTitle,
	}
}

func (c *setTitleCommand) Execute() {
	c.OldTitle = c.doc.GetTitle()
	c.doc.SetTitle(c.NewTitle)
}

func (c *setTitleCommand) Unexecute() {
	c.doc.SetTitle(c.OldTitle)
}

func (c *setTitleCommand) Merge(next Command) bool {
	nextCmd, ok := next.(*setTitleCommand)
	if ok {
		c.NewTitle = nextCmd.NewTitle
		return true
	}
	return false
}
