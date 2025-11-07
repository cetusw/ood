package mocks

import (
	"editor/pkg/document"
	"editor/pkg/model"

	"github.com/stretchr/testify/mock"
)

type MockDocument struct {
	mock.Mock
}

func (m *MockDocument) InsertParagraph(text string, position int) (document.Item, error) {
	args := m.Called(text, position)
	var item document.Item
	if args.Get(0) != nil {
		item = args.Get(0).(document.Item)
	}
	return item, args.Error(1)
}

func (m *MockDocument) InsertImage(path string, size model.Size, position int) (document.Item, error) {
	args := m.Called(path, size, position)
	var item document.Item
	if args.Get(0) != nil {
		item = args.Get(0).(document.Item)
	}
	return item, args.Error(1)
}

func (m *MockDocument) InsertItem(item document.Item, position int) (document.Item, error) {
	args := m.Called(item, position)
	var resItem document.Item
	if args.Get(0) != nil {
		resItem = args.Get(0).(document.Item)
	}
	return resItem, args.Error(1)
}

func (m *MockDocument) DeleteItem(index int) (document.Item, error) {
	args := m.Called(index)
	var item document.Item
	if args.Get(0) != nil {
		item = args.Get(0).(document.Item)
	}
	return item, args.Error(1)
}

func (m *MockDocument) Save(path string) error {
	args := m.Called(path)
	return args.Error(0)
}

func (m *MockDocument) GetItemsCount() int {
	args := m.Called()
	return args.Int(0)
}

func (m *MockDocument) GetItem(index int) (document.Item, error) {
	args := m.Called(index)
	var item document.Item
	if args.Get(0) != nil {
		item = args.Get(0).(document.Item)
	}
	return item, args.Error(1)
}

func (m *MockDocument) GetTitle() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockDocument) SetTitle(title string) {
	m.Called(title)
}

func (m *MockDocument) List() {
	m.Called()
}
