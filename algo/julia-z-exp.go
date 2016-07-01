package algo

import (
	"fmt"
	"math"
)

// Represents the iteration:  z * exp(z) + c
type juliazexp struct {
	depth int
	escape float64
	c     complex128
}

func (f *juliazexp) String() string {
	return fmt.Sprintf("Julia: x * exp(x) + %v (max depth %d)", f.c, f.depth)
}

func (f *juliazexp) ArgHelp() string {
	return "Arg1: real(c), Arg2: imag(c)."
}

func (f *juliazexp) Intensity(x, y float64) uint8 {
	val := complex(x, y)
	ans := f.depth - 1

	for ans = f.depth - 1; ans > 0; ans-- {
		if norm(val) >= f.escape {
			break
		}
		// e^(a+bi) == e^a(cos b + i sin b)
		ea := math.Exp(real(val))
		val = val * complex(ea*math.Cos(imag(val)),
			ea*math.Sin(imag(val))) + f.c
	}

	// scale the answer to the 0 - 255 range
	return uint8(float64(ans) * 256 / float64(f.depth))
}

// Returns a new Fractal which computes the
// Julia variation:  x*e^x + c.
func NewJuliaZExp(c complex128, depth int, escape float64) Fractal {
	return &juliazexp{c: c, depth: depth, escape: escape}
}
