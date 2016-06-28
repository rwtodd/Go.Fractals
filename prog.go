package main

import (
	"image"
	"image/color"
	"image/gif"
	"os"

	"github.com/rwtodd/fractals-go/algo"
)

const (
	picwidth  = 600
	picheight = 600
)

type fractalState struct {
	midx, midy, spanx, spany float64
}

func gray256() color.Palette {
	ans := make(color.Palette, 256)
	for i := range ans {
		ans[i] = color.Gray{uint8(i)}
	}
	return ans
}

func drawImage(fract algo.Fractal, fs *fractalState) image.Image {
	img := image.NewPaletted(image.Rectangle{image.Pt(0, 0),
		image.Pt(picwidth, picheight)},
		gray256())

	xUL := fs.midx - 0.5*fs.spanx
	yUL := fs.midy - 0.5*fs.spany
	for y := 0; y < picheight; y++ {
		ycoord := yUL + float64(y)*fs.spany/float64(picheight)
		for x := 0; x < picwidth; x++ {
			xcoord := xUL + float64(x)*fs.spanx/float64(picwidth)
			img.SetColorIndex(x, y, fract.Intensity(xcoord, ycoord))
		}
	}
	return img
}

func main() {
	f := algo.NewMandelbrot(256)
	fs := fractalState{midx: 0, midy: 0, spanx: 2, spany: 2}
	img := drawImage(f, &fs)
	gif.Encode(os.Stdout, img, nil)
}
