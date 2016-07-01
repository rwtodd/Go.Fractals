package algo

import (
	"fmt"
	"math"
)

type juliaexp struct {
	depth int
	escape float64 
	c     complex128
}

func (f *juliaexp) String() string {
	return fmt.Sprintf("Julia: exp(x) + %v (max depth %d)", f.c, f.depth)
}

func (f *juliaexp) ArgHelp() string {
	return "Arg1: real(c), Arg2: imag(c)."
}

func (f *juliaexp) Intensity(x, y float64) uint8 {
	val := complex(x, y)
	ans := f.depth - 1

	for ans = f.depth - 1; ans > 0; ans-- {
		if norm(val) >= f.escape {
			break
		}
		// e^(a+bi) == e^a(cos b + i sin b)
		ea := math.Exp(real(val))
		val = complex(ea*math.Cos(imag(val))+real(f.c),
			ea*math.Sin(imag(val))+imag(f.c))
	}

	// scale the answer to the 0 - 255 range
	return uint8(float64(ans) * 256 / float64(f.depth))
}

// Returns a new Fractal which computes the
// Julia variation:  e^x + c.
func NewJuliaExp(c complex128, depth int, escape float64) Fractal {
	return &juliaexp{c: c, depth: depth, escape: escape}
}
