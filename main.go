package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	const width int = 1919
	const height int = 1080
	var colours = [3]color.RGBA{Red, Green, Blue}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomImage := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < randomImage.Bounds().Max.Y; y++ {
		for x := 0; x < randomImage.Bounds().Max.X; x++ {
			randomImage.SetRGBA(x, y, colours[random.Intn(len(colours))])
		}
	}
	out, err := os.Create("output.png")
	if err != nil {
		log.Fatalln(err)
	}
	err = png.Encode(out, randomImage)
	if err != nil {
		log.Fatalln(err)
	}
	err = out.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

var (
	Red = color.RGBA{
		R: 255,
		G: 0,
		B: 0,
		A: 255,
	}
	Green = color.RGBA{
		R: 0,
		G: 255,
		B: 0,
		A: 255,
	}
	Blue = color.RGBA{
		R: 0,
		G: 0,
		B: 255,
		A: 255,
	}
)
