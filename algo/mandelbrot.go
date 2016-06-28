package algo

import (
	"fmt"
)

type mandel struct {
	depth int
}

func (f mandel) Name() string {
	return "Mandelbrot"
}

func (f mandel) String() string {
	return fmt.Sprintf("Mandelbrot (max depth %d)", f.depth)
}

func (f mandel) ArgHelp() string {
	return "No arguments."
}

func (f mandel) Intensity(x, y float64) uint8 {
	pos := complex(x, y)
	val := pos
	ans := f.depth - 1

	for ans = f.depth - 1; ans > 0; ans-- {
		if norm(val) >= 4.0 {
			break
		}
		val = val*val + pos
	}

	// scale the answer to the 0 - 255 range
	return uint8(float64(ans) * 256 / float64(f.depth))
}

// Returns a new Fractal which computes the
// standard mandelbrot set:  x^2 + location.
func NewMandelbrot(depth int) Fractal {
	return mandel{depth}
}
