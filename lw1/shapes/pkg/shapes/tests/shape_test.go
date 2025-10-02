package tests

import (
	"shapes/pkg/shapes/model"
	"shapes/pkg/shapes/shape"
	"testing"
)

func TestShapeNewAndGetters(t *testing.T) {
	mockStrategy := &MockStrategy{}

	sh := shape.NewShape(mockStrategy, "circ", "#000000")

	if sh.GetID() != "circ" {
		t.Errorf("GetID: Ожидался 'circ', получили '%s'", sh.GetID())
	}
	if sh.GetColor() != "#000000" {
		t.Errorf("GetColor: Ожидался '#000000', получили '%s'", sh.GetColor())
	}
	if sh.GetStrategy() != mockStrategy {
		t.Error("GetStrategy: Стратегия не была корректно установлена")
	}
}

func TestShapeMove(t *testing.T) {
	mockStrat := &MockStrategy{}
	sh := shape.NewShape(mockStrat, "circ", "#000000")

	moveVector := model.Point{X: 5.5, Y: -1.2}

	sh.GetStrategy().MoveShape(moveVector)

	if !mockStrat.MoveShapeCalled {
		t.Error("sh.GetStrategy().MoveShape не вызвал MoveShape на стратегии.")
	}
	if mockStrat.MoveVector.X != moveVector.X || mockStrat.MoveVector.Y != moveVector.Y {
		t.Errorf("Стратегия вызвана с неверным вектором. Ожидался %v, получен %v", moveVector, mockStrat.MoveVector)
	}
}

func TestShapeDraw(t *testing.T) {
	mockCanvas := &MockCanvas{}
	mockStrat := &MockStrategy{}
	sh := shape.NewShape(mockStrat, "circ", "#000000")

	sh.GetStrategy().Draw(mockCanvas, "#000000")

	if mockStrat.DrawCalled != true {
		t.Errorf("sh.GetStrategy().Draw не вызвал Draw на стратегии.")
	}
}
