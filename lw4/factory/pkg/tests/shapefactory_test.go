package tests

import (
	"errors"
	"strings"
	"testing"

	"factory/pkg/designer"
	"factory/pkg/tests/mocks"
)

func TestDesignerCreateDraft(t *testing.T) {
	mockFactory := &mocks.MockShapeFactory{
		ReturnShape: &mocks.MockShape{},
	}
	d := designer.NewDesigner(mockFactory)
	input := "rectangle red 10 20 30 40\nellipse blue 50 50 10 10"
	reader := strings.NewReader(input)

	draft, err := d.CreateDraft(reader)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if draft.GetShapeCount() != 2 {
		t.Errorf("Expected draft to have 2 shapes, but got %d", draft.GetShapeCount())
	}
	if mockFactory.ReceivedDescription != "ellipse blue 50 50 10 10" {
		t.Errorf("Expected factory to receive last description, but got '%s'", mockFactory.ReceivedDescription)
	}
}

func TestDesignerCreateDraftFailed(t *testing.T) {
	expectedErr := errors.New("factory failure")
	mockFactory := &mocks.MockShapeFactory{
		ReturnError: expectedErr,
	}
	d := designer.NewDesigner(mockFactory)
	input := "bad shape description"
	reader := strings.NewReader(input)

	_, err := d.CreateDraft(reader)

	if err == nil {
		t.Fatal("Expected designer to return an error, but it was nil")
	}
	if !errors.Is(err, expectedErr) {
		t.Errorf("Expected error to wrap factory error, but it did not")
	}
}
