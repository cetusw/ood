package tests

import (
	"testing"

	"factory/pkg/draft"
	"factory/pkg/tests/mocks"
)

func TestDrawDraft(t *testing.T) {
	shape1 := &mocks.MockShape{}
	shape2 := &mocks.MockShape{}
	d := draft.Draft{}
	d.AddShape(shape1)
	d.AddShape(shape2)
	mockCanvas := &mocks.MockCanvas{}

	d.Draw(mockCanvas)

	if !shape1.DrawCalled || !shape2.DrawCalled {
		t.Error("Expected Draw to be called on all shapes in the draft")
	}
}

func TestIncreaseShapeCount(t *testing.T) {
	d := draft.Draft{}

	d.AddShape(&mocks.MockShape{})
	d.AddShape(&mocks.MockShape{})

	if d.GetShapeCount() != 2 {
		t.Errorf("Expected shape count to be 2, but got %d", d.GetShapeCount())
	}
}
