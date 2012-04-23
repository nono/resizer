package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path"
	"strings"
)

type Ratio struct {
	w int
	h int
}

var dir string
var back image.Image
var ratios []Ratio

func box(xx int, yy int) image.Rectangle {
	rect := image.ZR
	surf := -1

	for _, ratio := range ratios {
		w, h := ratio.w, ratio.h
		x, y := xx, yy

		if (x-y)*(w-h) < 0 {
			w, h = h, w
		}

		if x*h > y*w {
			y = x * h / w
		} else {
			x = y * w / h
		}

		if x*y/surf < 1 {
			rect = image.Rect(0, 0, x, y)
			surf = x * y
		}
	}

	return rect
}

func resize(filename string) {
	in, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	src, format, err := image.Decode(in)
	if err != nil {
		log.Fatal(err)
	}

	rect := box(src.Bounds().Dx(), src.Bounds().Dy())
	dst := image.NewRGBA(rect)
	draw.Draw(dst, dst.Bounds(), back, image.ZP, draw.Src)
	draw.Draw(dst, src.Bounds(), src, image.ZP, draw.Src)

	file := path.Join(dir, path.Base(filename))
	out, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	switch format {
	case "png":
		png.Encode(out, dst)
	case "jpeg":
		jpeg.Encode(out, dst, nil)
	default:
		log.Fatal("Unknown format ", format)
	}
}

func main() {
	var r string
	var c string
	flag.StringVar(&dir, "dir", "resized", "Put the resized images in this directory")
	flag.StringVar(&r, "ratio", "4:3,3:2", "Use the best ratio from this list")
	flag.StringVar(&c, "color", "white", "Use this color for padding (white, black or transparent")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("No images to resize.\n")
	}

	switch c {
	case "white":
		back = image.White
	case "black":
		back = image.Black
	case "transparent":
		back = image.Transparent
	default:
		log.Fatal("Unknown color ", c)
	}
	parts := strings.Split(r, ",")
	ratios = make([]Ratio, len(parts))
	for i, part := range parts {
		fmt.Sscanf(part, "%d:%d", &ratios[i].w, &ratios[i].h)
	}
	os.MkdirAll(dir, 0755)

	for _, filename := range args {
		fmt.Printf("Resizing %s\n", filename)
		resize(filename)
	}
}
