package moderngraphicslib

import (
	"fmt"
	"io"

	"adapter/pkg/model"
)

type ModernGraphicsRenderer struct {
	out     io.Writer
	drawing bool
}

func NewModernGraphicsRenderer(out io.Writer) *ModernGraphicsRenderer {
	return &ModernGraphicsRenderer{out: out}
}

func (r *ModernGraphicsRenderer) BeginDraw() {
	if r.drawing {
		fmt.Println("Drawing has already begun")
		return
	}
	fmt.Fprintln(r.out, "<draw>")
	r.drawing = true
}

func (r *ModernGraphicsRenderer) DrawLine(start model.Point, end model.Point, color model.Color) {
	if !r.drawing {
		fmt.Println("DrawLine is allowed between BeginDraw()/EndDraw() only")
		return
	}

	fmt.Fprintf(
		r.out,
		"  <line fromX=\"%d\" fromY=\"%d\" toX=\"%d\" toY=\"%d\">\n    <color r=\"%.2f\" g=\"%.2f\" b=\"%.2f\" a=\"%.2f\"/>\n  </line>\n",
		start.X,
		start.Y,
		end.X,
		end.Y,
		color.R,
		color.G,
		color.B,
		color.A,
	)
}

func (r *ModernGraphicsRenderer) EndDraw() {
	if !r.drawing {
		fmt.Println("Drawing has not been started")
		return
	}
	fmt.Fprintln(r.out, "</draw>")
	r.drawing = false
}
