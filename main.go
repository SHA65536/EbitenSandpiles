package main

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	size  = 101
	scale = 8
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
	drawImage *image.RGBA
	sandPile  [size][size]int
}

func (g *Game) Update() error {
	var newPile [size][size]int
	var length = len(model)
	var stale = false
	g.sandPile[size/2][size/2] += 1
	for !stale {
		stale = true
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				newPile[i][j] = g.sandPile[i][j] % length
				for _, inst := range model {
					dx, dy := i+inst[0], j+inst[1]
					if dx < 0 || dy < 0 || dx >= size || dy >= size {
						continue
					}
					newPile[i][j] += g.sandPile[dx][dy] / length
				}
				if newPile[i][j] >= length {
					stale = false
				} else {
					g.drawImage.Set(i, j, colours[newPile[i][j]])
				}
			}
		}
		g.sandPile = newPile
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.ReplacePixels(g.drawImage.Pix)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return size, size
}

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
