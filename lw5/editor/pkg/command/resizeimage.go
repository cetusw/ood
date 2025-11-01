package command

import (
	"editor/pkg/document"
	"editor/pkg/model"
	"fmt"
)

type resizeImageCommand struct {
	AbstractCommand
	doc      document.Document
	position int
	newSize  model.Size
	oldSize  model.Size
	image    document.Image
}

func NewResizeImageCommand(doc document.Document, position int, newSize model.Size) Command {
	return &resizeImageCommand{
		doc:      doc,
		position: position,
		newSize:  newSize,
	}
}

func (c *resizeImageCommand) Execute() {
	if c.image != nil {
		c.image.Resize(c.newSize)
	}
	item, err := c.doc.GetItem(c.position)
	if err != nil {
		fmt.Println(err)
		return
	}
	img := item.(document.Image).GetImage()
	if img == nil {
		fmt.Printf("item at position %d is not image.\n", c.position+1)
		return
	}
	c.image = img
	c.image.Resize(c.newSize)
	c.oldSize = c.image.GetSize()
}

func (c *resizeImageCommand) Unexecute() {
	if c.image != nil {
		c.image.Resize(c.oldSize)
	}
}

func (c *resizeImageCommand) Merge(next Command) bool {
	nextCmd, ok := next.(*resizeImageCommand)
	if ok {
		c.newSize = nextCmd.newSize
		return true
	}
	return false
}
