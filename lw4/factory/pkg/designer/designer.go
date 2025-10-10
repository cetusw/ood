package designer

import (
	"bufio"
	"fmt"
	"io"

	"factory/pkg/draft"
	"factory/pkg/shapefactory"
)

type Designer interface {
	CreateDraft(reader io.Reader) (draft.Draft, error)
}

type designer struct {
	factory shapefactory.ShapeFactory
}

func NewDesigner(factory shapefactory.ShapeFactory) Designer {
	return &designer{factory: factory}
}

func (d *designer) CreateDraft(reader io.Reader) (draft.Draft, error) {
	pictureDraft := draft.Draft{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		shape, err := d.factory.CreateShape(line)
		if err != nil {
			return draft.Draft{}, fmt.Errorf("line '%s': %w", line, err)
		}
		pictureDraft.AddShape(shape)
	}
	return pictureDraft, scanner.Err()
}
