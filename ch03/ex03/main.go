// Copyright 2017 Ken Miura
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
			z1, ok1 := heightAt(i+1, j)
			z2, ok2 := heightAt(i, j)
			z3, ok3 := heightAt(i, j+1)
			z4, ok4 := heightAt(i+1, j+1)
			if !(ok1 && ok2 && ok3 && ok4) {
				continue
			}
			min = math.Min((z1+z2+z3+z4)/4, min)
			max = math.Max((z1+z2+z3+z4)/4, max)
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

			averageHeight := (heightAtA + heightAtB + heightAtC + heightAtD) / 4
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=%s/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, colorFormatStringBasedOnHeight(averageHeight, max, min))
		}
	}
	fmt.Println("</svg>")
}

func colorFormatStringBasedOnHeight(height, max, min float64) string {
	width := math.Abs(max - min)
	factor := int(math.Abs(height-min) * 255 / width)
	return fmt.Sprintf("'#%02x00%02x'", factor, 255-factor)
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
