package main

import (
	"adapter/pkg/model"
	"bufio"
	"fmt"
	"os"
	"strings"

	"adapter/pkg/adapter"
	"adapter/pkg/graphicslib"
	"adapter/pkg/moderngraphicslib"
	"adapter/pkg/shapedrawinglib"
)

func paintPicture(painter shapedrawinglib.CanvasPainter) {
	triangle := shapedrawinglib.NewTriangle(
		model.Point{X: 10, Y: 15},
		model.Point{X: 100, Y: 200},
		model.Point{X: 150, Y: 250},
	)
	rectangle := shapedrawinglib.NewRectangle(
		model.Point{X: 30, Y: 40}, 18, 24,
	)

	painter.Draw(triangle)
	painter.Draw(rectangle)
}

func paintPictureOnCanvas() {
	fmt.Println("Painting with old graphics library...")
	canvas := graphicslib.NewCanvas()
	canvas.SetColor(0xFF0000FF)
	painter := shapedrawinglib.NewCanvasPainter(canvas)
	paintPicture(painter)
}

func paintPictureOnModernGraphicsRenderer() {
	fmt.Println("Painting with modern graphics library...")
	renderer := moderngraphicslib.NewModernGraphicsRenderer(os.Stdout)
	a := adapter.NewModernRendererAdapter(renderer)
	a.SetColor(0xFF0000FF)

	renderer.BeginDraw()
	defer renderer.EndDraw()

	painter := shapedrawinglib.NewCanvasPainter(a)
	paintPicture(painter)
}

func main() {
	fmt.Print("Should we use new API (y)? ")
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)

	if strings.ToLower(userInput) == "y" {
		paintPictureOnModernGraphicsRenderer()
	} else {
		paintPictureOnCanvas()
	}
}
