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
	"runtime"
	"strings"
	"sync"
)

type Ratio struct {
	w int
	h int
}

var dir string
var back image.Image
var ratios []Ratio

// box function creates to complementary rectangle for the given image size `xx` and `yy`
// BUG(box): we always generate an image for the first ratio, not memory efficient
func box(xx int, yy int) image.Rectangle {
	rect := image.ZR
	surf := -1

	for _, ratio := range ratios {
		w, h := ratio.w, ratio.h
		x, y := xx, yy

		// img & ratio orientation fit
		if (x-y)*(w-h) < 0 {
			w, h = h, w
		}

		// complement on the right side
		if x*h > y*w {
			y = x * h / w
		} else {
			x = y * w / h
		}

		// select the best available ratio
		if x*y/surf < 1 {
			rect = image.Rect(0, 0, x, y)
			surf = x * y
		}
	}

	return rect
}

// resize function concatenates the given image with its complementary "bleed"
func resize(filename string, wg *sync.WaitGroup, runningjobs chan string) {
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
	<-runningjobs
	wg.Done()
}

// main triggers and waits for resizing goroutines
func main() {
	var wg sync.WaitGroup
	var r string
	var c string
	var p int

	flag.StringVar(&dir, "dir", "resized", "Put the resized images in this directory.")
	flag.StringVar(&r, "ratio", "4:3,3:2,5:4", "Use the best ratio from this list.")
	flag.StringVar(&c, "color", "white", "Use this color for padding (white, black or transparent.")
	flag.IntVar(&p, "parallel", runtime.NumCPU(), "Handle images in parallel, defaults to the actuel number of CPU available.")
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("No images to resize.\n")
	}
	
	runtime.GOMAXPROCS(p)
	var runningjobs = make(chan string, p)

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
		wg.Add(1)
		runningjobs <- filename
		fmt.Println("TODO: %s", filename)
		go resize(filename, &wg, runningjobs)
	}

	wg.Wait()
}
