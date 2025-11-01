package command

import (
	"editor/pkg/document"
	"fmt"
)

type deleteItemCommand struct {
	AbstractCommand
	doc         document.Document
	position    int
	deletedItem document.Item
}

func NewDeleteItemCommand(doc document.Document, position int) Command {
	return &deleteItemCommand{
		doc:      doc,
		position: position,
	}
}

func (c *deleteItemCommand) Execute() {
	item, err := c.doc.DeleteItem(c.position)
	if err != nil {
		fmt.Printf("Failed to execute DeleteItem: %v\n", err)
		return
	}
	c.deletedItem = item
}

func (c *deleteItemCommand) Unexecute() {
	_, err := c.doc.InsertItem(c.deletedItem, c.position)
	if err != nil {
		fmt.Printf("Failed to unexecute DeleteItem: %v\n", err)
		return
	}
	c.deletedItem = nil
}
