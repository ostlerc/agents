package main

import (
	"flag"
	"log"
	"os"

	"gopkg.in/qml.v1"
)

var (
	grid   Grid
	dialog qml.Object
	file   = flag.String("file", "", "json map to load at startup")
)

func init() {
	flag.Parse()
}

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
	grid.StatusText = win.Root().ObjectByName("statusText")
	grid.DefFoodCnt = win.Root().ObjectByName("defaultFoodCountCombo")
	grid.FoodCnt = win.Root().ObjectByName("countSpinner")
	grid.LifeCnt = win.Root().ObjectByName("lifeSpinner")
	grid.FoodTime = win.Root().ObjectByName("defaultFoodLifetimeCombo")
	grid.TileComp = &Tile{Object: tileComponent}

	dialog = win.Root().ObjectByName("fileDialog")

	if *file != "" {
		grid.LoadGrid(*file)
	} else {
		grid.BuildGrid()
	}

	win.Show()
	win.Wait()

	return nil
}
