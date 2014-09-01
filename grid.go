package main

import (
	"fmt"
	"math"

	"gopkg.in/qml.v1"
)

type Grid struct {
	Rows        qml.Object
	Cols        qml.Object
	Grid        qml.Object
	Tile        *Tile
	RunBtn      qml.Object
	Tiles       []*Tile
	Edited      bool
	ColumnCount int
	RowCount    int
	Start       *Tile
	End         *Tile
}

type Tile struct {
	Object qml.Object
}

func (t *Tile) index() int {
	return t.Object.Int("index")
}

func (t *Tile) Obj() qml.Object {
	return t.Object
}

func (t *Tile) Pos() (float64, float64) {
	i := t.index()
	x := float64(i % grid.ColumnCount)
	y := float64(i / grid.ColumnCount)
	return x, y
}

func (t *Tile) Neighbors() []Node {
	nodes := make([]Node, 0, 4)
	i := t.index()
	top := i - grid.ColumnCount
	if top >= 0 {
		nodes = append(nodes, grid.Tiles[top])
	}
	mod := i % grid.ColumnCount
	if mod != grid.ColumnCount-1 { //not end of column
		nodes = append(nodes, grid.Tiles[i+1])
	}
	if mod != 0 && i > 0 { //not beginning ofcolumn
		nodes = append(nodes, grid.Tiles[i-1])
	}
	bottom := i + grid.ColumnCount
	if bottom < len(grid.Tiles) {
		nodes = append(nodes, grid.Tiles[bottom])
	}
	return nodes
}

func (t *Tile) Dist(n Node) float64 {
	if t.Object.Int("type") == 1 {
		return 1e10
	}
	return t.EstimatedCost(n)
}

func (t *Tile) EstimatedCost(g Node) float64 {
	tx, ty := t.Pos()
	gx, gy := g.Pos()
	return math.Abs(tx-gx) + math.Abs(ty-gy)
}

func (g *Grid) SetStart(i int) {
	g.Start = g.Tiles[i]
	g.RunBtn.Set("enabled", g.Start != nil && g.End != nil)
}

func (g *Grid) SetEnd(i int) {
	g.End = g.Tiles[i]
	g.RunBtn.Set("enabled", g.Start != nil && g.End != nil)
}

func (g *Grid) ClearStart() {
	g.Start = nil
	g.RunBtn.Set("enabled", false)
}

func (g *Grid) ClearEnd() {
	g.End = nil
	g.RunBtn.Set("enabled", false)
}

func (g *Grid) createTile() *Tile {
	tile := &Tile{Object: g.Tile.Object.Create(nil)}
	tile.Object.Set("parent", g.Grid)
	return tile
}

func (g *Grid) index(col, row int) int {
	return col + (row * g.ColumnCount)
}

func (g *Grid) BuildGrid() {
	g.Edited = true
	for _, b := range g.Tiles {
		if b != nil {
			b.Object.Destroy()
		}
	}
	g.Start = nil
	g.End = nil
	g.RowCount = g.Rows.Int("value")
	g.ColumnCount = g.Cols.Int("value")
	g.Grid.Set("columns", g.ColumnCount)
	g.RunBtn.Set("enabled", false)

	fmt.Println("Building a", g.RowCount, g.ColumnCount, "grid")
	size := g.RowCount * g.ColumnCount
	g.Tiles = make([]*Tile, size, size)
	for n := 0; n < size; n++ {
		g.Tiles[n] = g.createTile()
		g.Tiles[n].Object.Set("index", n)
	}
}

func (g *Grid) RunAStar() {
	nodes := make([]Node, len(g.Tiles), len(g.Tiles))
	for i, v := range g.Tiles {
		nodes[i] = v
	}
	graph := NewAstar(nodes)
	path, err := graph.CalculatePath(g.Start, g.End)
	if err != nil {
		panic(err)
	}
	fmt.Println("finished")
	g.colorSolution(path)
}

func (g *Grid) TileClicked() {
	g.Edited = true
	for _, v := range g.Tiles {
		v.Object.Set("solution", false)
	}
}

func (g *Grid) colorSolution(objs []Node) {
	for _, v := range objs {
		if v.Obj().Int("type") == 1 {
			continue
		}
		x, y := v.Pos()
		fmt.Println("coloring", x, y)
		v.Obj().Set("solution", true)
	}
}
