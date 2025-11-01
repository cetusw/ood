package command

import (
	"fmt"
	"log"

	"editor/pkg/document"
)

type replaceTextCommand struct {
	AbstractCommand
	doc       document.Document
	position  int
	newText   string
	oldText   string
	paragraph document.Paragraph
}

func NewReplaceTextCommand(doc document.Document, position int, newText string) Command {
	return &replaceTextCommand{
		doc:      doc,
		position: position,
		newText:  newText,
	}
}

func (c *replaceTextCommand) Execute() {
	if c.paragraph != nil {
		c.paragraph.SetText(c.newText)
		return
	}
	item, err := c.doc.GetItem(c.position)
	if err != nil {
		fmt.Println(err)
		return
	}

	p := item.(document.Paragraph).GetParagraph()
	log.Println(p)
	if p == nil {
		fmt.Printf("item at position %d is not paragraph.\n", c.position+1)
		return
	}
	c.paragraph = p
	c.paragraph.SetText(c.newText)
	c.oldText = c.paragraph.GetText()
}

func (c *replaceTextCommand) Unexecute() {
	if c.paragraph != nil {
		c.paragraph.SetText(c.oldText)
	}
}

func (c *replaceTextCommand) Merge(next Command) bool {
	nextCmd, ok := next.(*replaceTextCommand)
	if ok {
		c.newText = nextCmd.newText
		return true
	}
	return false
}
