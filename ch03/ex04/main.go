// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

const (
	cells   = 100         // number of grid cells
	xyrange = 30.0        // axis ranges (-xyrange..+xyrange)
	angle   = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		width, height := 0.0, 0.0 // canvas size in pixels
		color := ""
		for k, v := range r.Form {
			var err error
			if k == "width" {
				width, err = strconv.ParseFloat(v[0], 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "parsing variable width: %v\n", err)
				}
			} else if k == "height" {
				height, err = strconv.ParseFloat(v[0], 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "parsing variable height: %v\n", err)
				}
			} else if k == "color" {
				color = "#" + v[0]
			}
		}
		// パースが失敗したり、異常値をパラメータとして渡されたり、パラメータを何も渡されなかったりしたときは下記の値に設定
		if width <= 0 {
			width = 600
		}
		if height <= 0 {
			height = 320
		}
		if !checkColorStringFormat(color) {
			color = "#ffffff"
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		sinc(w, width, height, color)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func checkColorStringFormat(color string) bool {
	return regexp.MustCompile(`#[0-9a-fA-F][0-9a-fA-F][0-9a-fA-F][0-9a-fA-F][0-9a-fA-F][0-9a-fA-F]`).Match([]byte(color))
}

func sinc(out io.Writer, width, height float64, color string) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aIsOk := corner(i+1, j, width, height)
			bx, by, bIsOk := corner(i, j, width, height)
			cx, cy, cIsOk := corner(i, j+1, width, height)
			dx, dy, dIsOk := corner(i+1, j+1, width, height)
			if !(aIsOk && bIsOk && cIsOk && dIsOk) {
				continue
			}
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int, width, height float64) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, false
	}

	// pixels per x or y unit
	xyscale := width / 2 / xyrange
	// pixels per z unit
	zscale := height * 0.4

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
