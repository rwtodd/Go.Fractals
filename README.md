# Fractals-Go

This is a web-based version of my
JavaFX [JFXMandelbrot][1] project, written in Go!

## Screenshots

It looks at least as good as the JavaFX version, I think.

![example 1](screenshots/shot1.png)

![example 2](screenshots/shot2.png)

## Status

It has feature parity with the JavaFX version now. 

You can control:
 * The algorithm
 * The amount of recursion/colors (depth)
 * The width and height of the image  
 * The center and span of the image on the complex plane

And, clicking on the image itself centers it on your click, optionally
zooming it in or out.

## Obtaining/Running

You can `go get`:

    go get github.com/rwtodd/fractals-go

To run on a local machine, start it with:

    fractals-go -local 8001

... to start it listening on port 8001 (for example).  The index and css files both need
to be in the directory you are running from.

If you start it without arguments it will attempt to run as a fcgi script.


[1]: https://github.com/rwtodd/JFXMandelbrot
