package main

import (
  "fmt"

  "github.com/rwtodd/fractals-go/algo"
)


func main() {
  f := algo.NewMandelbrot(500)
  var x, y float64
  for y = -1 ; y < 1 ; y += 0.1 {
  for x = -1 ; x < 1 ; x += 0.1 {
     fmt.Printf("At %v, %v = %v\n", x, y, f.Intensity(x,y)) 
  }
     fmt.Printf("\n")
  }
}
