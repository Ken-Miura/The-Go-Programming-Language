// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math"
	"math/cmplx"
	"net/http"
	"os"
	"strconv"
)

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		x0, y0 := 0.0, 0.0
		var scale float64 = 1
		for k, v := range r.Form {
			var err error
			if k == "x" {
				x0, err = strconv.ParseFloat(v[0], 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "parsing variable x: %v\n", err)
				}
			} else if k == "y" {
				y0, err = strconv.ParseFloat(v[0], 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "parsing variable y: %v\n", err)
				}
			} else if k == "scale" {
				scale, err = strconv.ParseFloat(v[0], 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "parsing variable scale: %v\n", err)
				}
			}
		}

		if math.IsNaN(x0) || math.IsInf(x0, 0) || x0 < xmin || x0 > xmax {
			fmt.Fprintf(w, "x must be %d <= x <= %d\n", xmin, xmax)
			return
		}
		if math.IsNaN(y0) || math.IsInf(y0, 0) || y0 < ymin || y0 > ymax {
			fmt.Fprintf(w, "y must be %d <= y <= %d\n", ymin, ymax)
			return
		}
		if scale <= 0 {
			fmt.Fprintln(w, "scale must be more than 0")
			return
		}
		fractal(w, x0, y0, scale)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func fractal(out io.Writer, x0, y0 float64, scale float64) {

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := (float64(py)/height*(ymax-ymin) + ymin) / scale
		for px := 0; px < width; px++ {
			x := (float64(px)/width*(xmax-xmin) + xmin) / scale
			z := complex(x-x0, y-y0)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(out, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
