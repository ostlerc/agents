package main

import (
	"fmt"

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
	Home       *Tile
	StatusText qml.Object
}

func (g *Grid) SetHome(i int) {
	g.Home = g.Tiles[i]
	g.RunBtn.Set("enabled", g.Home != nil)
}

func (g *Grid) ClearHome() {
	g.Home = nil
	g.RunBtn.Set("enabled", false)
}

func (g *Grid) createTile() *Tile {
	tile := &Tile{
		Object:   g.TileComp.Object.Create(nil),
		diagonal: true,
	}
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
	g.Home = nil
	g.RowCount = g.Rows.Int("value")
	g.ColCount = g.Cols.Int("value")
	g.Grid.Set("columns", g.ColCount)
	g.RunBtn.Set("enabled", false)
	g.StatusText.Set("text", "Click the grid cells to make a Home, end, and walls.")

	fmt.Println("Building a", g.RowCount, g.ColCount, "grid")
	size := g.RowCount * g.ColCount
	g.Tiles = make([]*Tile, size, size)
	for n := 0; n < size; n++ {
		g.Tiles[n] = g.createTile()
		g.Tiles[n].Object.Set("index", n)
	}
}

func (g *Grid) SetStatus(s string) {
	g.StatusText.Set("text", s)
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
