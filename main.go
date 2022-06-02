package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	size  = 201
	scale = 3
)

func main() {
	ebiten.SetWindowSize(size*scale, size*scale)
	ebiten.SetWindowTitle("Sandpiles")
	ebiten.SetMaxTPS(200)
	g := &Game{
		drawImage: image.NewRGBA(image.Rect(0, 0, size, size)),
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
