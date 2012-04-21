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
)

var dir string
var back image.Image
var width int
var height int

func box(x int, y int) image.Rectangle {
	rotate := false
	if (x-y)*(width-height) < 0 {
		rotate = !rotate
		x, y = y, x
	}

	if x*height > y*width {
		y = x * height / width
	} else {
		x = y * width / height
	}

	if rotate {
		x, y = y, x
	}
	return image.Rect(0, 0, x, y)
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
	flag.StringVar(&r, "ratio", "40:60", "Use this ratio for images (height:width)")
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
	fmt.Sscanf(r, "%d:%d", &width, &height)
	os.MkdirAll(dir, 0755)

	for _, filename := range args {
		fmt.Printf("Resizing %s\n", filename)
		resize(filename)
	}
}
