package main

import (
	"fmt"

	"gopkg.in/qml.v1"
)

type Grid struct {
	Rows        qml.Object
	Cols        qml.Object
	Grid        qml.Object
	Tile        qml.Object
	RunBtn      qml.Object
	Tiles       []qml.Object
	Edited      bool
	ColumnCount int
	RowCount    int
	Start       qml.Object
	End         qml.Object
}

func (g *Grid) ClearStart() {
	g.Start = nil
	g.RunBtn.Set("enabled", false)
}

func (g *Grid) ClearEnd() {
	g.End = nil
	g.RunBtn.Set("enabled", false)
}

func (g *Grid) createTile() qml.Object {
	tile := g.Tile.Create(nil)
	tile.Set("parent", g.Grid)
	return tile
}

func (g *Grid) index(col, row int) int {
	return col + (row * g.ColumnCount)
}

func (g *Grid) BuildGrid() {
	g.Edited = true
	for _, b := range g.Tiles {
		if b != nil {
			b.Destroy()
		}
	}
	g.Start = nil
	g.End = nil
	g.RowCount = int(g.Rows.Property("value").(float64))
	g.ColumnCount = int(g.Cols.Property("value").(float64))
	g.Grid.Set("columns", g.ColumnCount)
	g.RunBtn.Set("enabled", false)

	fmt.Println("Building a", g.RowCount, g.ColumnCount, "grid")
	size := g.RowCount * g.ColumnCount
	g.Tiles = make([]qml.Object, size, size)
	for n := 0; n < size; n++ {
		g.Tiles[n] = g.createTile()
		g.Tiles[n].Set("index", n)
	}
}

func (g *Grid) RunAStar() {
	visited := make(map[qml.Object]bool)
	visited[g.Start] = true
	fmt.Println(g.Start.Property("index"))
}

func (g *Grid) RunManhattan() {
	i := g.Start.Int("index")
	e := g.End.Int("index")
	var sol []qml.Object
	for i != e {
		if i < e {
			i++
		} else {
			i--
		}
		sol = append(sol, g.Tiles[i])
	}

	g.Edited = false
	g.colorSolution(sol)
}

func (g *Grid) TileClicked() {
	g.Edited = true
	for _, v := range g.Tiles {
		v.Set("solution", false)
	}
}

func (g *Grid) colorSolution(objs []qml.Object) {
	for _, v := range objs {
		if v.Int("type") == 1 {
			continue
		}
		v.Set("solution", true)
	}
}
