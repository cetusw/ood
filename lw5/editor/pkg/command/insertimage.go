package command

import (
	"editor/pkg/document"
	"editor/pkg/model"
	"fmt"
)

type insertImageCommand struct {
	AbstractCommand
	doc          document.Document
	path         string
	size         model.Size
	position     int
	insertedItem document.Item
}

func NewInsertImageCommand(doc document.Document, path string, size model.Size, position int) Command {
	return &insertImageCommand{
		doc:      doc,
		path:     path,
		size:     size,
		position: position,
	}
}

func (c *insertImageCommand) Execute() {
	if c.insertedItem != nil {
		_, err := c.doc.InsertItem(c.insertedItem, c.position)
		if err != nil {
			fmt.Printf("Failed to execute InsertImage: %v\n", err)
			return
		}
		c.insertedItem = nil
		return
	}
	item, err := c.doc.InsertImage(c.path, c.size, c.position)
	if err != nil {
		fmt.Printf("Failed to execute InsertImage: %v\n", err)
	}
	c.insertedItem = item
}

func (c *insertImageCommand) Unexecute() {
	_, err := c.doc.DeleteItem(c.position)
	if err != nil {
		fmt.Printf("Failed to unexecute InsertImage: %v\n", err)
	}
}

func (c *insertImageCommand) Destroy() {
	if c.insertedItem != nil {
		img, ok := c.insertedItem.(document.Image)
		if ok {
			img.Destroy()
		}
	}
}
