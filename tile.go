package main

import "gopkg.in/qml.v1"

type Tile struct {
	Object   qml.Object
	diagonal bool
	x        int
	y        int
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

func (t *Tile) Mark(x, y int) {
	t.x, t.y = x, y
	t.Object.Set("pcount", t.x)
}

func (t *Tile) Sniff() (int, int) {
	return t.x, t.y
}

func (t *Tile) Food() int {
	f := t.Object.Int("type")
	c := t.Object.Int("count")
	if f != 3 {
		return 0
	}
	return c
}

func (t *Tile) Snatch() Food {
	f := t.Object.Int("type")
	c := t.Object.Int("count")
	if f == 3 && c > 0 { //Food
		t.Object.Set("count", c-1)
		return Food{t.Object.Int("life")}
	}
	return Food{0}
}

func (t *Tile) Enter(food bool) {
	c := t.Object.Int("antcount")
	t.Object.Set("antcount", c+1)
	t.Object.Set("solution", food)
}

func (t *Tile) Exit() {
	c := t.Object.Int("antcount")
	t.Object.Set("antcount", c-1)
	t.Object.Set("solution", false)
}

func (t *Tile) Drop(Food) {
	grid.FoodQty++
}

func (t *Tile) Neighbors() []AntNode {
	neighbors := make([]AntNode, 0, 8)
	//cost := 1.0
	i := t.Object.Int("index")

	add := func(i int) {
		if grid.Tiles[i].Object.Int("type") != 1 {
			neighbors = append(neighbors, grid.Tiles[i])
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
