package main

import (
	"fmt"
	"os"

	"factory/pkg/canvas"
	"factory/pkg/designer"
	"factory/pkg/shapefactory"
)

func main() {
	shapeFactory := shapefactory.NewShapeFactory()
	appDesigner := designer.NewDesigner(shapeFactory)

	pngFileCanvas := canvas.NewPngCanvas(800, 600)

	fmt.Println("Enter shape descriptions (e.g., 'rectangle red 100 100 400 300'). Press Ctrl+D to finish.")
	draft, err := appDesigner.CreateDraft(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	draft.Draw(pngFileCanvas)

	outputFile := "picture.png"
	err = pngFileCanvas.SaveToFile(outputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to save file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n--- PNG file '%s' has been generated successfully. ---\n", outputFile)
}
