package ui

// Color 32 bit color value in ARGB format.
type Color struct {
	alpha   int
	blue    int
	green   int
	opacity float32
	red     int
}

// NewColor returns a Color
func NewColor() *Color {
	return &Color{}
}

// Value final color value
func (s *Color) Value() int {
	return 0
}
