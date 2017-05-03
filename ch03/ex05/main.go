// Copyright 2017 Ken Mirua
// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	// 発散する（2を超える）速さ毎に色分けするための定数
	const divergence1 = 1
	const divergence2 = 3
	const divergence3 = 5

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			if n < divergence1 {
				return color.Gray{255 - contrast*n}
			} else if n < divergence2 {
				return color.RGBA{0, 0, 255 - contrast*n, 255}
			} else if n < divergence3 {
				return color.RGBA{0, 255 - contrast*n, 0, 255}
			} else {
				return color.RGBA{255 - contrast*n, 0, 0, 255}
			}
		}
	}
	return color.Black
}
