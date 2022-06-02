package main

func SolveMatrix(g *Game) {
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
}

func SolveMatrixSeek(g *Game) {
	var newPile [size][size]int
	var length, start, end int
	var stale = false
	length = len(model)
	g.sandPile[size/2][size/2] += 1
	for i := 0; i <= size/2; i++ {
		if g.sandPile[i][size/2] != 0 {
			start = i - 4
			end = size - i + 4
			break
		}
		if g.sandPile[i][(size/2)-1] != 0 {
			start = i - 4
			end = size - i + 4
			break
		}
		if g.sandPile[i][(size/2)+1] != 0 {
			start = i - 4
			end = size - i + 4
			break
		}
	}
	if start < 0 {
		start = 0
		end = size
	}
	for !stale {
		stale = true
		for i := start; i < end; i++ {
			for j := start; j < end; j++ {
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
	///g.drawImage.Set(start, end-1, color.RGBA{0, 0, 254, 255})
	//g.drawImage.Set(end-1, start, color.RGBA{0, 0, 254, 255})
}
