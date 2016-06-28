package algo

import (
	"fmt"
	"math"
)

type juliaexp struct {
	depth int
	c     complex128
}

func (f *juliaexp) Name() string {
	return "Julia [exp(x) + c]"
}

func (f *juliaexp) String() string {
	return fmt.Sprintf("Jula exp(x) + %v (max depth %d)", f.c, f.depth)
}

func (f *juliaexp) ArgHelp() string {
	return "Arg1: real(c), Arg2: imag(c)."
}

func (f *juliaexp) Intensity(x, y float64) uint8 {
	val := complex(x, y)
	ans := f.depth - 1
	fc := f.c // make a local copy to help the optimizer

	for ans = f.depth - 1; ans > 0; ans-- {
		if norm(val) >= 4.0 {
			break
		}
		// e^(a+bi) == e^a(cos b + i sin b)
		ea := math.Exp(real(val))
		val = complex(ea*math.Cos(imag(val))+real(fc),
			ea*math.Sin(imag(val))+imag(fc))
	}

	// scale the answer to the 0 - 255 range
	return uint8(float64(ans) * 256 / float64(f.depth))
}

func NewJuliaExp(c complex128, depth int) Fractal {
	return &juliaexp{c: c, depth: depth}
}
