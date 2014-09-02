package main

import (
	"fmt"
	"strconv"
	"time"

	"gopkg.in/qml.v1"
)

type Grid struct {
	Rows       qml.Object
	Cols       qml.Object
	Grid       qml.Object
	TileComp   *Tile
	RunBtn     qml.Object
	Tiles      []*Tile
	Edited     bool
	ColCount   int
	RowCount   int
	Start      *Tile
	End        *Tile
	StatusText qml.Object
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
	tile := &Tile{Object: g.TileComp.Object.Create(nil)}
	tile.Object.Set("parent", g.Grid)
	return tile
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
	g.ColCount = g.Cols.Int("value")
	g.Grid.Set("columns", g.ColCount)
	g.RunBtn.Set("enabled", false)
	g.StatusText.Set("text", "Click the grid cells to make a start, end, and walls.")

	fmt.Println("Building a", g.RowCount, g.ColCount, "grid")
	size := g.RowCount * g.ColCount
	g.Tiles = make([]*Tile, size, size)
	for n := 0; n < size; n++ {
		g.Tiles[n] = g.createTile()
		g.Tiles[n].Object.Set("index", n)
	}
}

func (g *Grid) RunAStar() {
	g.ClearGrid()
	nodes := make([]Node, len(g.Tiles), len(g.Tiles))
	for i, v := range g.Tiles {
		nodes[i] = v
	}
	graph := NewAstar(nodes)
	start := time.Now()
	path, err := graph.CalculatePath(g.Start, g.End)
	if err != nil {
		panic(err)
	}
	elapsed := time.Since(start)
	timeStr := "Took " + strconv.Itoa(int(elapsed.Nanoseconds()/1000)) + " microseconds"
	if len(path) == 0 {
		g.StatusText.Set("text", "No valid path. "+timeStr)
	} else {
		g.StatusText.Set("text", timeStr)
	}
	g.colorSolution(path)
}

func (g *Grid) ClearGrid() {
	g.Edited = true
	for _, v := range g.Tiles {
		v.Object.Set("solution", false)
	}
}

func (g *Grid) colorSolution(objs []Node) {
	for _, v := range objs {
		v.(*Tile).Object.Set("solution", true)
	}
}
