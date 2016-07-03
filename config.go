package main

import (
	"net/url"
	"strconv"

	"github.com/rwtodd/fractals-go/algo"
)

// configuration to send to the HTML/JS portion of the
// app, which keeps all the logic needed to add another
// fractal in the Go part, unless new parameters need
// to be added.

type fractalConfig struct {
	ID      string
	Display string
	Params  []string
}

type fractalFactory map[string]func(url.Values) algo.Fractal

var configurations = []fractalConfig{
	{"Mandelbrot", "Mandelbrot", []string{}},
	{"JuliaSq", "Julia: v^2 + c", []string{"creal"}},
	{"JuliaExp", "Julia: exp(v) + c", []string{"creal", "esc"}},
	{"JuliaZExp", "Julia: v*exp(v) + c", []string{"creal", "esc"}},
}

var factory = fractalFactory{
	"Mandelbrot": func(args url.Values) algo.Fractal {
		depth, _ := strconv.Atoi(getOrElse(args["depth"], "256"))
		return algo.NewMandelbrot(depth)
	},
	"JuliaSq": func(args url.Values) algo.Fractal {
		depth, _ := strconv.Atoi(getOrElse(args["depth"], "256"))
		creal, _ := strconv.ParseFloat(getOrElse(args["creal"], "0.1"), 64)
		cimag, _ := strconv.ParseFloat(getOrElse(args["cimag"], "0.1"), 64)
		return algo.NewJuliaSq(complex(creal, cimag), depth)
	},
	"JuliaExp": func(args url.Values) algo.Fractal {
		depth, _ := strconv.Atoi(getOrElse(args["depth"], "256"))
		creal, _ := strconv.ParseFloat(getOrElse(args["creal"], "0.1"), 64)
		cimag, _ := strconv.ParseFloat(getOrElse(args["cimag"], "0.1"), 64)
		esc, _ := strconv.ParseFloat(getOrElse(args["esc"], "4.0"), 64)
		return algo.NewJuliaExp(complex(creal, cimag), depth, esc)
	},
	"JuliaZExp": func(args url.Values) algo.Fractal {
		depth, _ := strconv.Atoi(getOrElse(args["depth"], "256"))
		creal, _ := strconv.ParseFloat(getOrElse(args["creal"], "0.1"), 64)
		cimag, _ := strconv.ParseFloat(getOrElse(args["cimag"], "0.1"), 64)
		esc, _ := strconv.ParseFloat(getOrElse(args["esc"], "4.0"), 64)
		return algo.NewJuliaZExp(complex(creal, cimag), depth, esc)
	}}
