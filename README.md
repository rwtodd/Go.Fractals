# Fractals-Go

This is (going to be) a web version of my
JavaFX [JFXMandelbrot][1] project. 

## Screenshots

It looks at least as good as the JavaFX version, I think.

![example 1](screenshots/shot1.png)

![example 2](screenshots/shot2.png)

## Status

It's starting to come together. It has a UI now.  The main thing it's missing
is the ability to click on the image to zoom in and out.  I need to investigate
what javascript offers me in that area.

Start it with:

    $ fractals-go -local 8001

... to start it listening on port 8001.  The index and css files both need
to be in the same directory as the executable.


[1]: https://github.com/rwtodd/JFXMandelbrot
