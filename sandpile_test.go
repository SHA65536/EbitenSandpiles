package main

import (
	"image"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkMatrixFull(b *testing.B) {
	g := &Game{
		drawImage: image.NewRGBA(image.Rect(0, 0, size, size)),
	}
	for i := 0; i < 10000; i++ {
		SolveMatrix(g)
		g.count++
	}
}

func BenchmarkMatrixSeek(b *testing.B) {
	g := &Game{
		drawImage: image.NewRGBA(image.Rect(0, 0, size, size)),
	}
	for i := 0; i < 10000; i++ {
		SolveMatrixSeek(g)
		g.count++
	}
}

func TestSolveMatrix(t *testing.T) {
	assert := assert.New(t)
	g1 := &Game{
		drawImage: image.NewRGBA(image.Rect(0, 0, size, size)),
	}
	g2 := &Game{
		drawImage: image.NewRGBA(image.Rect(0, 0, size, size)),
	}
	for i := 0; i < 10000; i++ {
		SolveMatrix(g1)
		SolveMatrixSeek(g2)
		g1.count++
		g2.count++
	}
	assert.Equal(g1.sandPile, g2.sandPile, drawMatrix(g1)+"==========\n"+drawMatrix(g2))
}

func drawMatrix(g *Game) string {
	var res string
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			res += strconv.Itoa(g.sandPile[i][j])
		}
		res += "\n"
	}
	return res
}
