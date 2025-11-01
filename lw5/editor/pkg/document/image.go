package document

import (
	"editor/pkg/model"
	"fmt"
	"html"
	"log"
	"os"
)

type Image interface {
	Item
	Resize(size model.Size)
	Destroy()
	GetPath() string
	GetSize() model.Size
	GetImage() Image
}

type image struct {
	size model.Size
	path string
}

func NewImage(size model.Size, path string) Image {
	return &image{
		size: size,
		path: path,
	}
}

func (i *image) ToHTML() string {
	return fmt.Sprintf(
		`<img src="%s" width="%d" height="%d" />`,
		html.EscapeString(i.path), i.size.Width, i.size.Height,
	)
}

func (i *image) ToString() string {
	return fmt.Sprintf("Image: %d %d %s", i.size.Width, i.size.Height, i.path)
}

func (i *image) Destroy() {
	err := os.Remove(i.path)
	if err != nil {
		log.Printf("Failed to delete image file %s: %v", i.path, err)
	}
}

func (i *image) GetItemType() string {
	return "image"
}

func (i *image) Resize(size model.Size) {
	i.size = size
}

func (i *image) GetPath() string {
	return i.path
}

func (i *image) GetSize() model.Size {
	return i.size
}

func (i *image) GetImage() Image {
	return i
}
