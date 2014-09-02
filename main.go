package main

import (
	"log"
	"os"

	"gopkg.in/qml.v1"
)

var grid Grid

func main() {
	if err := qml.Run(run); err != nil {
		log.Fatalf("error: %v\n", err)
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

	context := engine.Context()
	context.SetVar("grid", &grid)

	win := component.CreateWindow(nil)

	grid.Rows = win.Root().ObjectByName("rows")
	grid.Cols = win.Root().ObjectByName("cols")
	grid.Grid = win.Root().ObjectByName("grid")
	grid.RunBtn = win.Root().ObjectByName("runBtn")
	grid.Tile = &Tile{Object: tileComponent}
	grid.BuildGrid()

	win.Show()
	win.Wait()

	return nil
}
