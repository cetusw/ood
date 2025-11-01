package document

import (
	"fmt"
	"html"
)

type Paragraph interface {
	Item
	GetText() string
	SetText(text string)
	GetParagraph() Paragraph
}

type paragraph struct {
	text string
}

func NewParagraph(text string) Paragraph {
	return &paragraph{
		text: text,
	}
}

func (p *paragraph) ToHTML() string {
	return fmt.Sprintf("<p>%s</p>", html.EscapeString(p.text))
}

func (p *paragraph) ToString() string {
	return fmt.Sprintf("Paragraph: %s", p.text)
}

func (p *paragraph) GetItemType() string {
	return "paragraph"
}

func (p *paragraph) GetText() string {
	return p.text
}

func (p *paragraph) GetParagraph() Paragraph {
	return p
}

func (p *paragraph) SetText(text string) {
	p.text = text
}
