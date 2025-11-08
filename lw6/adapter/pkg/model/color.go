package model

type Color struct {
	R float64
	G float64
	B float64
	A float64
}

func Uint32ToColor(rgba uint32) Color {
	r := uint8((rgba >> 24) & 0xFF)
	g := uint8((rgba >> 16) & 0xFF)
	b := uint8((rgba >> 8) & 0xFF)
	a := uint8(rgba & 0xFF)

	return Color{
		R: float64(r) / 255.0,
		G: float64(g) / 255.0,
		B: float64(b) / 255.0,
		A: float64(a) / 255.0,
	}
}
