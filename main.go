package main

import (
	"encoding/json"
	"flag"
	"image/gif"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/rwtodd/Go.AppUtil/resource"
	"github.com/rwtodd/Go.Fractals/algo"
)

var port = flag.String("port", "8000", "serve pages on this localhost port")
var help bool

// rscBase is the root location of our resources
var rscBase resource.Locator

func main() {
	var err error
	flag.BoolVar(&help, "help", false, "print this usage information")
	flag.BoolVar(&help, "h", false, "print this usage information")

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(1)
	}
	rscBase = resource.NewPathLocator([]string{"."},
		filepath.Join("github.com", "rwtodd", "fractals-go", "ui"))

	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/fract.css", cssHandler)
	http.HandleFunc("/img", imgHandler)
	http.HandleFunc("/cfg", cfgHandler)

	if err = http.ListenAndServe("localhost:"+*port, nil); err != nil {
		log.Fatal(err)
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if index, err := rscBase.Path("index.html"); err == nil {
		http.ServeFile(w, r, index)
	} else {
		log.Fatal(err)
	}
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	if css, err := rscBase.Path("fract.css"); err == nil {
		http.ServeFile(w, r, css)
	} else {
		log.Fatal(err)
	}
}

func cfgHandler(w http.ResponseWriter, r *http.Request) {
	cfg, err := json.Marshal(configurations)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(cfg)
}

func getOrElse(lst []string, def string) string {
	if len(lst) > 0 {
		def = lst[0]
	}
	return def
}

func getFractal(args url.Values) (f algo.Fractal) {
	fractal := getOrElse(args["fractal"], "Mandelbrot")
	ffunc, ok := factory[fractal]
	if ok {
		f = ffunc(args)
	} else {
		// just default to standard mandelbrot
		f = algo.NewMandelbrot(256)
	}
	return
}

// imgHandler generates a fractal image
func imgHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	desiredWidth, _ := strconv.Atoi(getOrElse(r.Form["width"], "600"))
	desiredHeight, _ := strconv.Atoi(getOrElse(r.Form["height"], "600"))
	centerX, _ := strconv.ParseFloat(getOrElse(r.Form["cX"], "0.0"), 64)
	centerY, _ := strconv.ParseFloat(getOrElse(r.Form["cY"], "0.0"), 64)
	spanX, _ := strconv.ParseFloat(getOrElse(r.Form["spX"], "3.0"), 64)
	spanY, _ := strconv.ParseFloat(getOrElse(r.Form["spY"], "3.0"), 64)

	f := getFractal(r.Form)
	log.Print(f)
	fs := fractalState{
		midx:      centerX,
		midy:      centerY,
		spanx:     spanX,
		spany:     spanY,
		picwidth:  desiredWidth,
		picheight: desiredHeight,
	}
	img := drawImage(f, &fs)
	if err := gif.Encode(w, img, nil); err != nil {
		log.Print(err)
	}
}
