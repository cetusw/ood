package common

type Color struct {
	R float64
	G float64
	B float64
}

const PictureFileName = "picture.png"

const (
	Circle    = "circle"
	Rectangle = "rectangle"
	Triangle  = "triangle"
	Line      = "line"
	Text      = "text"
)

var (
	WhiteColor = Color{
		R: 1,
		G: 1,
		B: 1,
	}
)
