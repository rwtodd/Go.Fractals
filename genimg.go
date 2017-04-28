package main

import (
	"image"
	"image/color"
	"runtime"
	"sync"

	"github.com/rwtodd/Go.Fractals/algo"
)

type fractalState struct {
	picwidth, picheight      int
	midx, midy, spanx, spany float64
}

func gray256() color.Palette {
	ans := make(color.Palette, 256)
	for i := range ans {
		ans[i] = color.Gray{uint8(i)}
	}
	return ans
}

func drawFractal(fract algo.Fractal,
	fs *fractalState,
	img *image.Paletted,
	stY int,
	endY int) {
	xUL := fs.midx - 0.5*fs.spanx
	yUL := fs.midy + 0.5*fs.spany
	for y := stY; y < endY; y++ {
		ycoord := yUL - float64(y)/float64(fs.picheight)*fs.spany
		for x := 0; x < fs.picwidth; x++ {
			xcoord := xUL + float64(x)/float64(fs.picwidth)*fs.spanx
			img.SetColorIndex(x, y, fract.Intensity(xcoord, ycoord))
		}
	}
}

func drawImage(fract algo.Fractal, fs *fractalState) image.Image {
	var wg sync.WaitGroup
	img := image.NewPaletted(image.Rectangle{image.Pt(0, 0),
		image.Pt(fs.picwidth, fs.picheight)},
		gray256())

	divisions := runtime.NumCPU()
	divSize := fs.picheight / divisions

	wg.Add(divisions)
	for div := 0; div < divisions; div++ {
		go func(div int) {
			endLine := (div + 1) * divSize
			if div == (divisions - 1) {
				endLine = fs.picheight
			}
			drawFractal(fract, fs, img, div*divSize, endLine)
			wg.Done()
		}(div)
	}

	wg.Wait()
	return img
}
