package main

import (
	"fmt"
	"log"
	"slides/pkg/canvas"
	"slides/pkg/model"
	"slides/pkg/shapes"
	"slides/pkg/slide"
)

func main() {
	cnv := canvas.NewPngCanvas(800, 600)

	ground := createRect(0, 0, 800, 150, model.Green)
	ground.SetLineStyle(shapes.NewStyle(false, model.Undefined))

	sky := createRect(0, 150, 800, 450, model.Blue)
	sky.SetLineStyle(shapes.NewStyle(false, model.Undefined))

	sun := shapes.NewEllipse(model.Point{X: 700, Y: 500}, model.Radius{X: 40, Y: 40})
	sun.SetFillStyle(shapes.NewStyle(true, model.Yellow))
	sun.SetLineStyle(shapes.NewStyle(false, model.Undefined))

	houseGroup := shapes.NewGroup()
	houseBody := createRect(100, 150, 200, 150, model.Red)

	roof := shapes.NewPolygon([]model.Point{
		{X: 80, Y: 300},
		{X: 320, Y: 300},
		{X: 200, Y: 420},
	}, shapes.NewStyle(true, model.Black), shapes.NewStyle(true, model.Black))

	window := createRect(160, 200, 80, 60, model.Blue)
	window.SetLineStyle(shapes.NewStyle(true, model.Yellow))

	houseGroup.AddShape(houseBody)
	houseGroup.AddShape(roof)
	houseGroup.AddShape(window)

	log.Println("houseBody:", houseBody.GetFrame().X, houseBody.GetFrame().Y, houseBody.GetFrame().Width, houseBody.GetFrame().Height)
	log.Println("roof:", roof.GetFrame().X, roof.GetFrame().Y, roof.GetFrame().Width, roof.GetFrame().Height)
	log.Println("houseGroup:", houseGroup.GetFrame().X, houseGroup.GetFrame().Y, houseGroup.GetFrame().Width, houseGroup.GetFrame().Height)

	treeGroup := shapes.NewGroup()

	trunk := createRect(550, 150, 40, 120, model.Black)
	leaves := shapes.NewEllipse(model.Point{X: 610, Y: 320}, model.Radius{X: 40, Y: 40})
	leaves.SetFillStyle(shapes.NewStyle(true, model.Green))
	leaves.SetLineStyle(shapes.NewStyle(true, model.Black))

	treeGroup.AddShape(trunk)
	treeGroup.AddShape(leaves)

	scene := shapes.NewGroup()
	scene.AddShape(sky)
	scene.AddShape(ground)
	scene.AddShape(sun)
	scene.AddShape(houseGroup)
	scene.AddShape(treeGroup)

	s := slide.NewSlide(800, 600)
	s.InsertShape(scene)

	s.Draw(cnv)

	if err := cnv.SaveToFile("result.png"); err != nil {
		fmt.Printf("Ошибка сохранения: %v\n", err)
	} else {
		fmt.Println("Файл 'result.png' успешно создан.")
	}
}

func createRect(x, y, w, h float64, fillColor model.Color) shapes.Shape {
	p1 := model.Point{X: x, Y: y}
	p2 := model.Point{X: x + w, Y: y}
	p3 := model.Point{X: x + w, Y: y + h}
	p4 := model.Point{X: x, Y: y + h}

	lineStyle := shapes.NewStyle(true, model.Undefined)
	fillStyle := shapes.NewStyle(true, fillColor)

	return shapes.NewPolygon([]model.Point{p1, p2, p3, p4}, lineStyle, fillStyle)
}
