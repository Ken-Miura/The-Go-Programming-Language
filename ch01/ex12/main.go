// Copyright 2017 Ken Miura
// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		var cycles, res float64 = 0, 0
		var size, nframes, delay int = 0, 0, 0
		for k, v := range r.Form {
			var err error
			if k == "cycles" {
				cycles, err = strconv.ParseFloat(v[0], 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "parsing variable cycles: %v\n", err)
				}
			} else if k == "res" {
				res, err = strconv.ParseFloat(v[0], 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "parsing variable res: %v\n", err)
				}
			} else if k == "size" {
				size, err = strconv.Atoi(v[0])
				if err != nil {
					fmt.Fprintf(os.Stderr, "parsing variable size: %v\n", err)
				}
			} else if k == "nframes" {
				nframes, err = strconv.Atoi(v[0])
				if err != nil {
					fmt.Fprintf(os.Stderr, "parsing variable nframes: %v\n", err)
				}
			} else if k == "delay" {
				if err != nil {
					fmt.Fprintf(os.Stderr, "parsing variable delay: %v\n", err)
				}
				delay, err = strconv.Atoi(v[0])
			}
		}
		// パースが失敗したり、異常値をパラメータとして渡されたり、パラメータを何も渡されなかったりしたときは下記の値に設定
		if cycles <= 0 {
			cycles = 5
		}
		if res <= 0 {
			res = 0.001
		}
		if size <= 0 {
			size = 100
		}
		if nframes <= 0 {
			nframes = 64
		}
		if delay <= 0 {
			delay = 8
		}
		lissajous(w, cycles, res, size, nframes, delay)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer, cycles, res float64, size, nframes, delay int) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
