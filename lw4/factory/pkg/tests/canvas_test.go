package tests

import (
	"testing"

	"factory/pkg/canvas"
)

func TestNewPngCanvas(t *testing.T) {
	pngCanvas := canvas.NewPngCanvas(100, 100)

	if pngCanvas == nil {
		t.Fatal("NewPngCanvas returned nil")
	}
	if pngCanvas.C == nil {
		t.Error("Internal canvas object was not initialized")
	}
	if pngCanvas.Context == nil {
		t.Error("Internal context object was not initialized")
	}
}
