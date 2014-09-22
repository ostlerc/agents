package main

import (
	"log"
	"os"

	"gopkg.in/qml.v1"
)

var grid Grid
var dialog qml.Object

func main() {
	if err := qml.Run(run); err != nil {
		log.Fatalf("error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	engine := qml.NewEngine()

	component, err := engine.LoadFile("agents.qml")
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
	grid.StatusText = win.Root().ObjectByName("statusText")
	grid.DefFoodCnt = win.Root().ObjectByName("defaultFoodCountCombo")
	grid.FoodCnt = win.Root().ObjectByName("countSpinner")
	grid.LifeCnt = win.Root().ObjectByName("lifeSpinner")
	grid.FoodTime = win.Root().ObjectByName("defaultFoodLifetimeCombo")
	grid.TileComp = &Tile{Object: tileComponent}
	grid.BuildGrid()

	dialog = win.Root().ObjectByName("fileDialog")

	win.Show()
	win.Wait()

	return nil
}
