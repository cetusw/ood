package tests

import (
	"testing"

	"factory/pkg/model"
	"factory/pkg/shapes"
	"factory/pkg/tests/mocks"
)

func TestRectangleDraw(t *testing.T) {
	mockCanvas := &mocks.MockCanvas{}
	rect := shapes.NewRectangle(
		model.Red,
		model.Point{X: 10, Y: 20},
		model.Point{X: 50, Y: 40},
	)

	rect.Draw(mockCanvas)

	expectedLog := "SetColor(red); DrawLine(from: {10 20}, to: {50 20}); DrawLine(from: {50 20}, to: {50 40}); DrawLine(from: {50 40}, to: {10 40}); DrawLine(from: {10 40}, to: {10 20})"
	if mockCanvas.GetLog() != expectedLog {
		t.Errorf("Expected canvas calls '%s', but got '%s'", expectedLog, mockCanvas.GetLog())
	}
}

func TestTriangleDraw(t *testing.T) {
	mockCanvas := &mocks.MockCanvas{}
	tri := shapes.NewTriangle(
		model.Blue,
		model.Point{X: 10, Y: 10},
		model.Point{X: 20, Y: 30},
		model.Point{X: 0, Y: 30},
	)

	tri.Draw(mockCanvas)

	expectedLog := "SetColor(blue); DrawLine(from: {10 10}, to: {20 30}); DrawLine(from: {20 30}, to: {0 30}); DrawLine(from: {0 30}, to: {10 10})"
	if mockCanvas.GetLog() != expectedLog {
		t.Errorf("Expected canvas calls '%s', but got '%s'", expectedLog, mockCanvas.GetLog())
	}
}

func TestEllipseDraw(t *testing.T) {
	mockCanvas := &mocks.MockCanvas{}
	ellipse := shapes.NewEllipse(
		model.Green,
		model.Point{X: 100, Y: 150},
		model.Radius{X: 30, Y: 40},
	)

	ellipse.Draw(mockCanvas)

	expectedLog := "SetColor(green); DrawEllipse(center: {100 150}, radius.X: 30.00, radius.Y: 40.00)"
	if mockCanvas.GetLog() != expectedLog {
		t.Errorf("Expected canvas calls '%s', but got '%s'", expectedLog, mockCanvas.GetLog())
	}
}

func TestPolygonDraw(t *testing.T) {
	mockCanvas := &mocks.MockCanvas{}
	polygon := shapes.NewPolygon(
		model.Yellow,
		model.Point{X: 100, Y: 100},
		10,
		4,
	)

	polygon.Draw(mockCanvas)

	expectedLog := "SetColor(yellow); DrawLine(from: {110 100}, to: {100 110}); DrawLine(from: {100 110}, to: {90 100}); DrawLine(from: {90 100}, to: {100 90}); DrawLine(from: {100 90}, to: {110 100})"
	if mockCanvas.GetLog() != expectedLog {
		t.Errorf("Expected canvas calls '%s', but got '%s'", expectedLog, mockCanvas.GetLog())
	}
}
