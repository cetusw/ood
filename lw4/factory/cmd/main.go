package main

import (
	"fmt"
	"os"

	"factory/pkg/canvas"
	"factory/pkg/designer"
	"factory/pkg/factory"
)

func main() {
	f := factory.NewShapeFactory()
	c := canvas.NewCanvas()
	d := designer.NewDesigner(f)

	fmt.Println("Enter shape descriptions:")

	draft, err := d.CreateDraft(os.Stdin)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\n--- Drawing Picture Draft ---")
	draft.Draw(c)
	fmt.Println("--- Drawing Finished      ---")
}
