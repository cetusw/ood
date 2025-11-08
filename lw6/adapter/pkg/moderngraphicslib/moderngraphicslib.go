package moderngraphicslib

import (
	"adapter/pkg/model"
	"fmt"
	"io"
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
		panic("Drawing has already begun")
	}
	fmt.Fprintln(r.out, "<draw>")
	r.drawing = true
}

func (r *ModernGraphicsRenderer) DrawLine(start, end model.Point) {
	if !r.drawing {
		panic("DrawLine is allowed between BeginDraw()/EndDraw() only")
	}

	fmt.Fprintf(
		r.out,
		"  <line fromX=\"%d\" fromY=\"%d\" toX=\"%d\" toY=\"%d\"/>\n",
		start.X,
		start.Y,
		end.X,
		end.Y,
	)
}

func (r *ModernGraphicsRenderer) EndDraw() {
	if !r.drawing {
		panic("Drawing has not been started")
	}
	fmt.Fprintln(r.out, "</draw>")
	r.drawing = false
}
