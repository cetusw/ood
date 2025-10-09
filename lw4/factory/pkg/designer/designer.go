package designer

import (
	"bufio"
	"fmt"
	"io"

	"factory/pkg/domain"
	"factory/pkg/factory"
)

type Designer interface {
	CreateDraft(reader io.Reader) (domain.PictureDraft, error)
}

type designer struct {
	factory factory.ShapeFactory
}

func NewDesigner(factory factory.ShapeFactory) Designer {
	return &designer{factory: factory}
}

func (d *designer) CreateDraft(reader io.Reader) (domain.PictureDraft, error) {
	draft := domain.PictureDraft{}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		shape, err := d.factory.CreateShape(line)
		if err != nil {
			return domain.PictureDraft{}, fmt.Errorf("line '%s': %w", line, err)
		}
		draft.AddShape(shape)
	}
	return draft, scanner.Err()
}
