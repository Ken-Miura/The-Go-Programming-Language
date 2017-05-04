// Copyright 2017 Ken Mirua
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"runtime"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float32(px)/width*(xmax-xmin) + xmin
			z := complex(float64(x), float64(y))
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(complex64(z)))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
	memstats := runtime.MemStats{}
	runtime.ReadMemStats(&memstats)
	fmt.Fprintf(os.Stderr, "cumulative bytes allocated for heap objects : %d byte(s)\n", memstats.TotalAlloc)
}

func mandelbrot(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
