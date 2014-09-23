package main

import "gopkg.in/qml.v1"

type Tile struct {
	Object   qml.Object
	diagonal bool
}

type JSONTile struct {
	Type  int `json:"type,omitempty"`
	Count int `json:"count,omitempty"`
	Life  int `json:"life,omitempty"`
	Index int `json:"index,omitempty"`
}

//TODO: maybe remove this
func (t *Tile) Pos() (int, int) {
	i := t.Object.Int("index")
	x := i % grid.ColCount
	y := i / grid.ColCount
	return x, y
}

func (t *Tile) Mark(int, int) {
}

func (t *Tile) Sniff() (int, int) {
	return 0, 0
}

func (t *Tile) Snatch() Food {
	return Food{-1}
}

func (t *Tile) Enter() {
	typ := t.Object.Int("type")
	if typ != 0 && typ != 4 { //open or ant
		return //no counting necessary
	}

	c := t.Object.Int("count")
	t.Object.Set("count", c+1)
	t.Object.Set("type", 4) //ant
}

func (t *Tile) Exit() {
	typ := t.Object.Int("type")
	if typ != 0 && typ != 4 { //open or ant
		return //no counting necessary
	}

	c := t.Object.Int("count")
	t.Object.Set("count", c+1)
	if c == 0 {
		t.Object.Set("type", 0) //open
	} else {
		t.Object.Set("type", 4) //ant
	}
}

func (t *Tile) Drop(Food) {
}

func (t *Tile) Neighbors() []AntNode {
	neighbors := make([]AntNode, 0, 8)
	//cost := 1.0
	i := t.Object.Int("index")

	add := func(i int) {
		if grid.Tiles[i].Object.Int("type") != 1 {
			neighbors = append(neighbors, &grid.Tiles[i])
			//t.neighbors[&grid.Tiles[i]] = cost
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
