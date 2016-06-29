package algo

import (
	"fmt"
)

type juliasq struct {
	depth int
	c     complex128
}

func (f *juliasq) String() string {
	return fmt.Sprintf("Jula x^2 + %v (max depth %d)", f.c, f.depth)
}

func (f *juliasq) ArgHelp() string {
	return "Arg1: real(c), Arg2: imag(c)."
}

func (f *juliasq) Intensity(x, y float64) uint8 {
	val := complex(x, y)
	ans := f.depth - 1
	fc := f.c // make a local copy to help the optimizer

	for ans = f.depth - 1; ans > 0; ans-- {
		if norm(val) >= 4.0 {
			break
		}
		val = val*val + fc
	}

	// scale the answer to the 0 - 255 range
	return uint8(float64(ans) * 256 / float64(f.depth))
}

// Returns a new Fractal which computes the 
// Julia variation:  x^2 + c.
func NewJuliaSq(c complex128, depth int) Fractal {
	return &juliasq{c: c, depth: depth}
}
