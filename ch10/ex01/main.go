// Copyright 2017 Ken Miura
package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var format = flag.String("format", "jpeg", "output image format (gif, jpeg or png)")

func main() {
	flag.Parse()
	if !(*format == "gif" || *format == "jpeg" || *format == "png") {
		fmt.Fprintf(os.Stderr, "Unsupported image format: %v\n", *format)
		return
	}
	if err := toSpecifiedFormatImage(os.Stdin, os.Stdout, *format); err != nil {
		fmt.Fprintf(os.Stderr, "image formatter: %v\n", err)
		os.Exit(1)
	}
}

func toSpecifiedFormatImage(in io.Reader, out io.Writer, format string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch format {
	case "gif":
		return toGIF(out, img)
	case "jpeg":
		return toJPEG(out, img)
	case "png":
		return toPNG(out, img)
	default:
		panic(fmt.Sprintf("unsupported image format: %v", format))
	}
}

func toJPEG(out io.Writer, img image.Image) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(out io.Writer, img image.Image) error {
	return png.Encode(out, img)
}

func toGIF(out io.Writer, img image.Image) error {
	return gif.Encode(out, img, nil)
}
