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

// TODO bugがありそう
func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		raty := big.NewRat(0, 1)
		raty.SetFloat64(y)
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			ratx := big.NewRat(0, 1)
			ratx.SetFloat64(x)
			z := NewComplex(ratx, raty)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
	memstats := runtime.MemStats{}
	runtime.ReadMemStats(&memstats)
	fmt.Fprintf(os.Stderr, "cumulative bytes allocated for heap objects : %d byte(s)\n", memstats.TotalAlloc)
}

func mandelbrot(z *Complex) color.Color {
	// プログラムが終了しないので繰り返し回数少なめで
	const iterations = 3
	const contrast = 15

	four := big.NewRat(4, 1)

	v := NewComplex(big.NewRat(0, 1), big.NewRat(0, 1))
	for n := uint8(0); n < iterations; n++ {
		temp := Multiply(v, v)
		v = Add(temp, z)
		ret := SquaredAbs(v)
		if ret.Cmp(four) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
