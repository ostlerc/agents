package main

import (
	"math"

	"gopkg.in/qml.v1"
)

type Tile struct {
	Object qml.Object
}

func (t *Tile) index() int {
	return t.Object.Int("index")
}

func (t *Tile) Pos() (float64, float64) {
	i := t.index()
	x := float64(i % grid.ColCount)
	y := float64(i / grid.ColCount)
	return x, y
}

func (t *Tile) Neighbors() []Node {
	nodes := make([]Node, 0, 4)
	i := t.index()
	top := i - grid.ColCount
	if top >= 0 {
		if grid.Tiles[top].Object.Int("type") != 1 {
			nodes = append(nodes, grid.Tiles[top])
		}
	}
	mod := i % grid.ColCount
	if mod != grid.ColCount-1 { //not end of column
		if grid.Tiles[i+1].Object.Int("type") != 1 {
			nodes = append(nodes, grid.Tiles[i+1])
		}
	}
	if mod != 0 && i > 0 { //not beginning ofcolumn
		if grid.Tiles[i-1].Object.Int("type") != 1 {
			nodes = append(nodes, grid.Tiles[i-1])
		}
	}
	bottom := i + grid.ColCount
	if bottom < len(grid.Tiles) {
		if grid.Tiles[bottom].Object.Int("type") != 1 {
			nodes = append(nodes, grid.Tiles[bottom])
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
	return math.Abs(tx-gx) + math.Abs(ty-gy)
}
