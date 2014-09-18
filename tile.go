package main

import (
	"math"

	"gopkg.in/qml.v1"
)

type Tile struct {
	Object   qml.Object
	diagonal bool
}

func (t *Tile) Pos() (float64, float64) {
	i := t.Object.Int("index")
	x := float64(i % grid.ColCount)
	y := float64(i / grid.ColCount)
	return x, y
}

func (t *Tile) Neighbors() []Node {
	nodes := make([]Node, 0, 8)
	i := t.Object.Int("index")

	add := func(i int) {
		if grid.Tiles[i].Object.Int("type") != 1 {
			nodes = append(nodes, grid.Tiles[i])
		}
	}

	top := i - grid.ColCount
	if top >= 0 { //Up
		add(top)
	}
	mod := i % grid.ColCount
	if mod != grid.ColCount-1 { //Right
		add(i + 1)
	}
	if mod != 0 && i > 0 { //Left
		add(i - 1)
	}
	bottom := i + grid.ColCount
	if bottom < len(grid.Tiles) { //Bottom
		add(bottom)
	}
	if t.diagonal {
		if top >= 0 { //Top
			if mod != grid.ColCount-1 { //right
				add(i + 1 - grid.ColCount)
			}
			if mod != 0 && i > 0 { //left
				add(i - 1 - grid.ColCount)
			}
		}
		if bottom < len(grid.Tiles) { //Bottom
			if mod != grid.ColCount-1 { //right
				add(i + 1 + grid.ColCount)
			}
			if mod != 0 && i > 0 { //left
				add(i - 1 + grid.ColCount)
			}
		}
	}
	return nodes
}

func (t *Tile) Dist(n Node) float64 {
	return 1
}

func (t *Tile) EstimatedCost(g Node) float64 {
	tx, ty := t.Pos()
	gx, gy := g.(*Tile).Pos()
	if t.diagonal {
		return math.Max(math.Abs(tx-gx), math.Abs(ty-gy))
	}
	return math.Abs(tx-gx) + math.Abs(ty-gy)
}
