package main

import (
	"fmt"
	"os"

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
	HasStart    bool
	HasEnd      bool
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
	for _, b := range g.Tiles {
		if b != nil {
			b.Destroy()
		}
	}
	g.HasStart = false
	g.HasEnd = false
	g.Edited = true
	g.RowCount = int(g.Rows.Property("value").(float64))
	g.ColumnCount = int(g.Cols.Property("value").(float64))
	g.Grid.Set("columns", g.ColumnCount)
	g.RunBtn.Set("enabled", false)

	fmt.Println("Building a", g.RowCount, g.ColumnCount, "grid")
	size := g.RowCount * g.ColumnCount
	g.Tiles = make([]qml.Object, size, size)
	for n := 0; n < size; n++ {
		g.Tiles[n] = g.createTile()
	}
}

func (g *Grid) RunAStar() {
}

func main() {
	if err := qml.Run(run); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	engine := qml.NewEngine()

	component, err := engine.LoadFile("astar.qml")
	if err != nil {
		return err
	}

	tileComponent, err := engine.LoadFile("Tile.qml")
	if err != nil {
		return err
	}

	grid := Grid{}

	context := engine.Context()
	context.SetVar("grid", &grid)

	win := component.CreateWindow(nil)

	grid.Rows = win.Root().ObjectByName("rows")
	grid.Cols = win.Root().ObjectByName("cols")
	grid.Grid = win.Root().ObjectByName("grid")
	grid.RunBtn = win.Root().ObjectByName("runBtn")
	grid.Tile = tileComponent

	win.Show()
	win.Wait()

	return nil
}
