package main

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/qml.v1"
)

type Grid struct {
	Rows        qml.Object
	Cols        qml.Object
	Grid        qml.Object
	ColumnCount int
	RowCount    int
	HasStart    bool
	HasEnd      bool
}

func (g *Grid) BuildGrid() {
}

func (g *Grid) RunAStar() {
}

func (g *Grid) RowsClicked() {
	g.RowCount, _ = strconv.Atoi((g.Rows.Property("text")).(string))
}

func (g *Grid) ColsClicked() {
	g.ColumnCount, _ = strconv.Atoi((g.Cols.Property("text")).(string))
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

	grid := Grid{HasStart: true, HasEnd: true, ColumnCount: 10, RowCount: 10}

	context := engine.Context()
	context.SetVar("grid", &grid)

	win := component.CreateWindow(nil)

	grid.Rows = win.Root().ObjectByName("rows")
	grid.Cols = win.Root().ObjectByName("cols")
	grid.Grid = win.Root().ObjectByName("grid")

	win.Show()
	win.Wait()

	return nil
}
