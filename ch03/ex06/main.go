// Copyright 2017 Ken Miura
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: " + os.Args[0] + " 'input file' 'output file'")
		fmt.Println("ex. " + os.Args[0] + " input.png output.png")
		return
	}

	inputFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("failed to open file. error: %v\n", err)
		return
	}
	defer inputFile.Close()

	inputImg, err := png.Decode(inputFile)
	if err != nil {
		fmt.Printf("failed to decode file as png. error: %v\n", err)
		return
	}

	outputImg := average(inputImg)
	outputFile, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Printf("failed to create file. error: %v\n", err)
		return
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, outputImg)
	if err != nil {
		fmt.Printf("failed to encode image. error: %v\n", err)
		return
	}
}

// 画像の任意のピクセルに対して上下左右の4画素の値を取得し、平均化した画像を返す。
func average(img image.Image) image.Image {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	averagedImg := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			center := img.At(x, y)
			north := img.At(x, checkMaxAndCorrect(y+1, height))
			east := img.At(checkMaxAndCorrect(x+1, width), y)
			west := img.At(checkMinAndCorrect(x-1, 0), y)
			south := img.At(x, checkMinAndCorrect(y-1, 0))

			cr, cg, cb, ca := center.RGBA()
			nr, ng, nb, na := north.RGBA()
			er, eg, eb, ea := east.RGBA()
			wr, wg, wb, wa := west.RGBA()
			sr, sg, sb, sa := south.RGBA()

			r := (cr>>8 + nr>>8 + er>>8 + wr>>8 + sr>>8) / 5
			g := (cg>>8 + ng>>8 + eg>>8 + wg>>8 + sg>>8) / 5
			b := (cb>>8 + nb>>8 + eb>>8 + wb>>8 + sb>>8) / 5
			a := (ca>>8 + na>>8 + ea>>8 + wa>>8 + sa>>8) / 5

			averagedImg.Set(x, y, color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)})
		}
	}
	return averagedImg
}

func checkMinAndCorrect(value, min int) int {
	if value < min {
		return min
	} else {
		return value
	}
}

func checkMaxAndCorrect(value, max int) int {
	if value > max {
		return max
	} else {
		return value
	}
}
