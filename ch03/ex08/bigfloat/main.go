// Copyright 2017 Ken Mirua
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
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
		y := float64(py)/height*(ymax-ymin) + ymin
		bigy := big.NewFloat(53)
		bigy.SetFloat64(y)
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			bigx := big.NewFloat(53)
			bigx.SetFloat64(x)
			z, _ := NewBigComplex(bigx, bigy)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
	memstats := runtime.MemStats{}
	runtime.ReadMemStats(&memstats)
	fmt.Fprintf(os.Stderr, "cumulative bytes allocated for heap objects : %d byte(s)\n", memstats.TotalAlloc)
}

func mandelbrot(z *BigComplex) color.Color {
	const iterations = 10
	const contrast = 15

	two := big.NewFloat(2.0)

	v, _ := NewBigComplex(big.NewFloat(0.0), big.NewFloat(0.0))
	for n := uint8(0); n < iterations; n++ {
		temp := Multiply(v, v)
		v = Add(temp, z)
		ret := Abs(v)
		if ret.Cmp(two) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
