package tests

import (
	"log"
	"shapes/pkg/shapes"
	"shapes/pkg/shapes/model"
	"shapes/pkg/shapes/shape"
	"testing"
)

func TestPictureAddShape(t *testing.T) {
	pic := shapes.NewPicture()

	sh := shape.NewShape(&MockStrategy{}, "sh", "#000000")

	if err := pic.AddShape(sh); err != nil {
		t.Fatalf("AddShape вернул ошибку: %v", err)
	}

	if pic.GetShapes()[0] != sh {
		t.Error("Добавленная фигура в срезе не соответствует исходному объекту")
	}

	err := pic.AddShape(sh)
	if err == nil {
		t.Error("Ожидалась ошибка при добавлении фигуры с дублирующимся ID")
	}
}

func TestPictureMoveShape(t *testing.T) {
	pic := shapes.NewPicture()

	mockStrategy := &MockStrategy{}
	sh := shape.NewShape(mockStrategy, "sh", "#000000")

	if err := pic.AddShape(sh); err != nil {
		t.Fatalf("AddShape с ошибкой: %v", err)
	}

	moveVector := model.Point{X: 10, Y: 20}

	pic.MoveShape("sh", moveVector)

	if !mockStrategy.MoveShapeCalled {
		t.Error("MoveShape не вызвал MoveShape на стратегии фигуры.")
	}
	if mockStrategy.MoveVector.X != moveVector.X || mockStrategy.MoveVector.Y != moveVector.Y {
		t.Errorf("Стратегия вызвана с неверным вектором. Ожидался %v, получен %v", moveVector, mockStrategy.MoveVector)
	}
}

func TestPictureMovePicture(t *testing.T) {
	pic := shapes.NewPicture()

	mockStrategy1 := &MockStrategy{}
	mockStrategy2 := &MockStrategy{}

	sh1 := shape.NewShape(mockStrategy1, "sh1", "#000000")
	sh2 := shape.NewShape(mockStrategy2, "sh2", "#000000")

	if err := pic.AddShape(sh1); err != nil {
		t.Fatalf("AddShape с ошибкой: %v", err)
	}
	if err := pic.AddShape(sh2); err != nil {
		t.Fatalf("AddShape с ошибкой: %v", err)
	}

	moveVector := model.Point{X: 5, Y: 5}

	pic.MovePicture(moveVector)

	if !mockStrategy1.MoveShapeCalled || !mockStrategy2.MoveShapeCalled {
		t.Error("MovePicture не вызвал MoveShape на всех стратегиях.")
	}

	if mockStrategy1.MoveVector != moveVector {
		t.Errorf("Передан неверный вектор: %v", mockStrategy1.MoveVector)
	}
}

func TestPictureDeleteShape(t *testing.T) {
	pic := shapes.NewPicture()

	mockStrategy1 := &MockStrategy{}
	mockStrategy2 := &MockStrategy{}
	mockStrategy3 := &MockStrategy{}

	sh1 := shape.NewShape(mockStrategy1, "sh1", "#000000")
	sh2 := shape.NewShape(mockStrategy2, "sh2", "#000000")
	sh3 := shape.NewShape(mockStrategy3, "sh3", "#000000")

	if err := pic.AddShape(sh1); err != nil {
		t.Fatalf("AddShape с ошибкой: %v", err)
	}
	if err := pic.AddShape(sh2); err != nil {
		t.Fatalf("AddShape с ошибкой: %v", err)
	}
	if err := pic.AddShape(sh3); err != nil {
		t.Fatalf("AddShape с ошибкой: %v", err)
	}

	initialLength := len(pic.GetShapes())
	if initialLength != 3 {
		t.Fatalf("Ожидалось 3 фигуры, получено %d", initialLength)
	}

	pic.DeleteShape("sh2")

	log.Println(pic.GetShapes())

	if len(pic.GetShapes()) != 2 {
		t.Errorf("Длина среза не изменилась. Ожидалось 2, получено %d", len(pic.GetShapes()))
	}

	shs := pic.GetShapes()
	foundS2 := false
	for _, s := range shs {
		if s.GetID() == "sh2" {
			foundS2 = true
		}
	}
	if foundS2 {
		t.Error("Фигура 'sh2' не была удалена.")
	}

	pic.DeleteShape("nonexistent")
	if len(pic.GetShapes()) != 2 {
		t.Error("Длина среза изменилась после попытки удаления несуществующей фигуры.")
	}
}

func TestPictureChangeColor(t *testing.T) {
	pic := shapes.NewPicture()

	mockStrategy := &MockStrategy{}
	sh := shape.NewShape(mockStrategy, "sh", "#000000")
	if err := pic.AddShape(sh); err != nil {
		t.Fatalf("AddShape с ошибкой: %v", err)
	}

	newColor := "#FF0000"

	pic.ChangeColor("sh", newColor)

	if sh.GetColor() != newColor {
		t.Errorf("Цвет не изменился. Ожидался %s, получен %s", newColor, sh.GetColor())
	}
}

func TestPictureChangeShape(t *testing.T) {
	pic := shapes.NewPicture()

	initialStrat := &MockStrategy{}
	sh := shape.NewShape(initialStrat, "sh", "#000000")
	if err := pic.AddShape(sh); err != nil {
		t.Fatalf("AddShape с ошибкой: %v", err)
	}

	newStrat := &MockStrategy{}

	pic.ChangeShape("sh", newStrat)

	if sh.GetStrategy() != newStrat {
		t.Error("Стратегия фигуры не была изменена.")
	}
}

func TestPictureDrawShape(t *testing.T) {
	mockCanvas := &MockCanvas{}
	pic := shapes.NewPicture()

	mockStrategy := &MockStrategy{}
	sh := shape.NewShape(mockStrategy, "sh", "#000000")
	if err := pic.AddShape(sh); err != nil {
		t.Fatalf("AddShape с ошибкой: %v", err)
	}

	pic.DrawShape("sh", mockCanvas)

	if !mockStrategy.DrawCalled {
		t.Error("DrawShape не вызвал Draw на стратегии фигуры 's1'.")
	}
}

func TestPictureDrawPicture(t *testing.T) {
	mockCanvas := &MockCanvas{}
	pic := shapes.NewPicture()

	mockStrategy1 := &MockStrategy{}
	mockStrategy2 := &MockStrategy{}

	sh1 := shape.NewShape(mockStrategy1, "sh1", "#000000")
	sh2 := shape.NewShape(mockStrategy2, "sh2", "#000000")

	if err := pic.AddShape(sh1); err != nil {
		t.Fatalf("AddShape с ошибкой: %v", err)
	}
	if err := pic.AddShape(sh2); err != nil {
		t.Fatalf("AddShape с ошибкой: %v", err)
	}

	pic.DrawPicture(mockCanvas)

	if !mockStrategy1.DrawCalled || !mockStrategy2.DrawCalled {
		t.Error("DrawPicture не вызвал Draw на всех стратегиях.")
	}
}
