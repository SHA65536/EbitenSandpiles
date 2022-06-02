package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	Knight    = [][2]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {2, -1}, {2, 1}, {1, -2}, {1, 2}}
	Adajacent = [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
)

var colours = map[int]color.Color{
	0: color.Black,
	1: color.RGBA{212, 253, 4, 255},
	2: color.RGBA{253, 253, 4, 255},
	3: color.RGBA{251, 212, 0, 255},
	4: color.RGBA{251, 158, 0, 255},
	5: color.RGBA{249, 124, 0, 255},
	6: color.RGBA{254, 63, 2, 255},
	7: color.RGBA{254, 0, 0, 255},
}

var model [][2]int = Knight

type Game struct {
	count     int
	drawImage *image.RGBA
	sandPile  [size][size]int
}

func (g *Game) Update() error {
	g.count++
	SolveMatrixSeek(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(g.drawImage.Pix)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return size, size
}
