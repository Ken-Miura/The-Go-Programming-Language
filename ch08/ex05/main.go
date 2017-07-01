// Copyright 2017 Ken Miura
// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
	"time"
)

var isParallel = flag.Bool("parallel", false, "whether program executes mandelbrot in parallel or not")

func main() {
	flag.Parse()

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	if *isParallel {
		processInParallel(img, height, width, xmin, xmax, ymin, ymax)
	} else {
		processSequentially(img, height, width, xmin, xmax, ymin, ymax)
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func processSequentially(img *image.RGBA, height, width int, xmin, xmax, ymin, ymax float64) {
	start := time.Now()
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	fmt.Fprintf(os.Stderr, "time for creating mandelbrot image sequentially: %s", time.Since(start))
}

// サンプルコードのfloat64を使ったmandelbrot関数だと並列に実施した方が逐次処理より遅かった。
// TODO より処理の重い関数で実施
// TODO 使うゴルーチンの最適な数？？？
func processInParallel(img *image.RGBA, height, width int, xmin, xmax, ymin, ymax float64) {
	type colorPoint struct {
		x int
		y int
		c color.Color
	}
	ch := make(chan colorPoint, height*width)

	start := time.Now()
	var wg sync.WaitGroup
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			wg.Add(1)
			go func(px, py int) {
				defer wg.Done()
				z := complex(x, y)
				// Image point (px, py) represents complex value z.
				ch <- colorPoint{px, py, mandelbrot(z)}
			}(px, py)
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for cp := range ch {
		img.Set(cp.x, cp.y, cp.c)
	}
	fmt.Fprintf(os.Stderr, "time for creating mandelbrot image in parallel: %s", time.Since(start))
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

//!-
