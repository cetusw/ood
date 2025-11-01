package document

import (
	"fmt"
	"html"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"editor/pkg/model"

	"github.com/google/uuid"
)

const (
	defaultRights = 0755
	imagesDirName = "images"
)

type Document interface {
	InsertParagraph(text string, position int) (Item, error)
	InsertImage(path string, size model.Size, position int) (Item, error)
	InsertItem(item Item, position int) (Item, error)
	DeleteItem(index int) (Item, error)
	Save(path string) error
	GetItemsCount() int
	GetItem(index int) (Item, error)
	GetTitle() string
	SetTitle(title string)
	List()
}

type document struct {
	title         string
	documentItems []Item
}

func NewDocument(title string) Document {
	return &document{
		title:         title,
		documentItems: make([]Item, 0),
	}
}

func (d *document) InsertParagraph(text string, position int) (Item, error) {
	if position < 0 || position > len(d.documentItems) {
		return nil, fmt.Errorf("invalid position %d for insertion", position)
	}
	p := NewParagraph(text)
	d.insertItemAt(p, position)
	return p, nil
}

func (d *document) InsertImage(path string, size model.Size, position int) (Item, error) {
	if position < 0 || position > len(d.documentItems) {
		return nil, fmt.Errorf("invalid position %d for insertion", position)
	}

	imagesDir := imagesDirName
	newPath, err := copyAndGeneratePath(path, imagesDir)
	if err != nil {
		return nil, fmt.Errorf("failed to copy image: %w", err)
	}

	img := NewImage(size, newPath)
	d.insertItemAt(img, position)
	return img, nil
}

func (d *document) InsertItem(item Item, position int) (Item, error) {
	if position < 0 || position > len(d.documentItems) {
		return nil, fmt.Errorf("invalid position %d for insertion", position)
	}
	d.insertItemAt(item, position)
	return item, nil
}

func (d *document) GetItemsCount() int {
	return len(d.documentItems)
}

func (d *document) GetItem(index int) (Item, error) {
	if index < 0 || index >= len(d.documentItems) {
		return nil, fmt.Errorf("index %d out of range", index)
	}
	return d.documentItems[index], nil
}

func (d *document) DeleteItem(index int) (Item, error) {
	if index < 0 || index >= len(d.documentItems) {
		return nil, fmt.Errorf("index %d out of range", index)
	}
	item := d.documentItems[index]
	d.documentItems = append(d.documentItems[:index], d.documentItems[index+1:]...)
	return item, nil
}

func (d *document) GetTitle() string {
	return d.title
}

func (d *document) SetTitle(title string) {
	d.title = title
}

func (d *document) Save(path string) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(filepath.Join(dir, imagesDirName), defaultRights)
	if err != nil {
		return fmt.Errorf("failed to create directory %s: %w", imagesDirName, err)
	}

	var sb strings.Builder
	sb.WriteString("<!DOCTYPE html>\n<html>\n<head>\n")
	sb.WriteString(fmt.Sprintf("<title>%s</title>\n", html.EscapeString(d.title)))
	sb.WriteString("</head>\n<body>\n")

	for _, item := range d.documentItems {
		sb.WriteString(item.ToHTML() + "\n")
	}

	sb.WriteString("</body>\n</html>")

	return os.WriteFile(path, []byte(sb.String()), defaultRights)
}

func (d *document) List() {
	log.Printf("Title: %s", d.title)
	for i, item := range d.documentItems {
		log.Printf("%d. %s", i+1, item.ToString())
	}
}

func (d *document) insertItemAt(item Item, position int) {
	if position == len(d.documentItems) {
		d.documentItems = append(d.documentItems, item)
		return
	}
	d.documentItems = append(d.documentItems, nil)
	copy(d.documentItems[position+1:], d.documentItems[position:])
	d.documentItems[position] = item
}

func copyAndGeneratePath(srcPath, destDir string) (string, error) {
	err := os.MkdirAll(destDir, defaultRights)
	if err != nil {
		return "", err
	}
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return "", err
	}
	defer srcFile.Close()

	ext := filepath.Ext(srcPath)
	newName := uuid.New().String() + ext
	destPath := filepath.Join(destDir, newName)

	destFile, err := os.Create(destPath)
	if err != nil {
		return "", err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	return destPath, err
}
