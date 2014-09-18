package main

import (
	"math"

	"gopkg.in/qml.v1"
)

type Tile struct {
	Object     qml.Object
	diagonal   bool
	dneighbors map[Node]float64
}

func (t *Tile) Pos() (float64, float64) {
	i := t.Object.Int("index")
	x := float64(i % grid.ColCount)
	y := float64(i / grid.ColCount)
	return x, y
}

func (t *Tile) Neighbors() []Node {
	neighbors := make([]Node, 0, 8)
	t.dneighbors = make(map[Node]float64)
	cost := 1.0
	i := t.Object.Int("index")

	add := func(i int) {
		if grid.Tiles[i].Object.Int("type") != 1 {
			neighbors = append(neighbors, grid.Tiles[i])
			t.dneighbors[grid.Tiles[i]] = cost
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
		//playing with diagonal cost.
		//cost = math.sqrt(2)
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
	return neighbors
}

func (t *Tile) Dist(n Node) float64 {
	if v, ok := t.dneighbors[n]; ok {
		return v
	}
	panic("Incorrect non neighbor node")
}

func (t *Tile) EstimatedCost(g Node) float64 {
	tx, ty := t.Pos()
	gx, gy := g.(*Tile).Pos()
	if t.diagonal {
		return math.Max(math.Abs(tx-gx), math.Abs(ty-gy))
	}
	return math.Abs(tx-gx) + math.Abs(ty-gy)
}
