package main

import "gopkg.in/qml.v1"

type Node interface {
	Neighbors() []Node
}

type Tile struct {
	Object    qml.Object
	diagonal  bool
	neighbors map[Node]float64
}

type JSONTile struct {
	Type  int `json:"type,omitempty"`
	Count int `json:"count,omitempty"`
	Life  int `json:"life,omitempty"`
	Index int `json:"index,omitempty"`
}

//TODO: maybe remove this
func (t *Tile) Pos() (float64, float64) {
	i := t.Object.Int("index")
	x := float64(i % grid.ColCount)
	y := float64(i / grid.ColCount)
	return x, y
}

func (t *Tile) Neighbors() []Node {
	neighbors := make([]Node, 0, 8)
	t.neighbors = make(map[Node]float64)
	cost := 1.0
	i := t.Object.Int("index")

	add := func(i int) {
		if grid.Tiles[i].Object.Int("type") != 1 {
			neighbors = append(neighbors, &grid.Tiles[i])
			t.neighbors[&grid.Tiles[i]] = cost
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
