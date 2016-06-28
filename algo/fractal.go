// Package fractals-go/algo contains the algorithms
// for computing the pixel values of various fractal
// images.
package algo

import (
  "fmt"
)

// Fractal is the interface that decouples the driver program
// from the particular fractal being rendered.
type Fractal interface {
   fmt.Stringer

   // Intensity determines the pixel value of the given
   // coordinates.
   Intensity(x,y float64) uint8

   // Name gives the name of the fractal equation/algorithm.
   Name() string
 
   // ArgHelp provides a help string for use in the UI,
   // telling the user what the arguments to that fractal
   // represent.
   ArgHelp() string

}


