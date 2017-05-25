package main

import (
	"github.com/jonbonazza/mapgen"
	"image"
	"os"
	"image/png"
	"log"
	"math/rand"
	"flag"
	"time"
	"math"
)

func main() {
	var it int
	var points int
	var seed int64
	var width float64
	var height float64
	var unit float64
	flag.IntVar(&it, "it", 1, "")
	flag.IntVar(&points, "points", 4096, "")
	flag.Int64Var(&seed, "seed", time.Now().Unix(), "")
	flag.Float64Var(&width, "width", 512, "")
	flag.Float64Var(&height, "height", 512, "")
	flag.Float64Var(&height, "unit", 0, "")
	flag.Parse()
	rand.Seed(seed)
	if unit <= 0 {
		unit = math.Min(width, height)/20.0
	}
	bbox := mapgen.NewBBox(0, width, 0, height)
	m := mapgen.NewMap(bbox, points, it, unit)
	img := m.Image()
	if err := writeImgToFile(img, "out.png"); err != nil {
		log.Fatal(err)
	}
}

func writeImgToFile(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, img)
}
