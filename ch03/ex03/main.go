// Copyright 2017 Ken Mirua
// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	max, min := -math.MaxFloat64, math.MaxFloat64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			z, ok := heightAt(i, j)
			if !ok {
				continue
			}
			min = math.Min(z, min)
			max = math.Max(z, max)
		}
	}

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, aIsOk, heightAtA := corner(i+1, j)
			bx, by, bIsOk, heightAtB := corner(i, j)
			cx, cy, cIsOk, heightAtC := corner(i, j+1)
			dx, dy, dIsOk, heightAtD := corner(i+1, j+1)
			if !(aIsOk && bIsOk && cIsOk && dIsOk) {
				continue
			}

			tolerance := 0.05
			if (math.Abs(max-heightAtA) < tolerance) || (math.Abs(max-heightAtB) < tolerance) || (math.Abs(max-heightAtC) < tolerance) || (math.Abs(max-heightAtD) < tolerance) {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#ff0000'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			} else if (math.Abs(min-heightAtA) < tolerance) || (math.Abs(min-heightAtB) < tolerance) || (math.Abs(min-heightAtC) < tolerance) || (math.Abs(min-heightAtD) < tolerance) {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#0000ff'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			} else {
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func heightAt(i, j int) (float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, false
	}
	return z, true
}

func corner(i, j int) (float64, float64, bool, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, false, 0
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, true, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
