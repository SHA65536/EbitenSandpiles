# Ebiten Sandpiles
This repository is an example of using [Ebitengine](https://ebiten.org/) to draw 'sandpiles' on screen.

## Algorithm
Each tick we increase the middle of the matrix by 1, once a cell in the matrix goes over the model threshold, the cell 'topples' to the surrounding cells and empties itself.

## Video
[![Sample Video](https://github.com/SHA65536/EbitenSandpiles/blob/main/.giuthub/Sample.png?raw=true)](https://streamable.com/5kjlil)