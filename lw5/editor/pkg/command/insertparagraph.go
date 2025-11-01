package command

import (
	"editor/pkg/document"
	"log"
)

type insertParagraphCommand struct {
	AbstractCommand
	doc          document.Document
	position     int
	text         string
	insertedItem document.Item
}

func NewInsertParagraphCommand(doc document.Document, position int, text string) Command {
	return &insertParagraphCommand{
		doc:      doc,
		text:     text,
		position: position,
	}
}

func (c *insertParagraphCommand) Execute() {
	item, err := c.doc.InsertParagraph(c.text, c.position)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	c.insertedItem = item
}

func (c *insertParagraphCommand) Unexecute() {
	_, err := c.doc.DeleteItem(c.position)
	if err != nil {
		log.Printf("Error: %s", err)
	}
}
