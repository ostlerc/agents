package main

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/qml.v1"
)

type Grid struct {
	Text qml.Object
}

func (g *Grid) Clicked() {
	s := (g.Text.Property("text")).(string)
	i, _ := strconv.Atoi(s)
	g.Text.Set("text", strconv.Itoa(i+1))
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

	grid := Grid{}

	context := engine.Context()
	context.SetVar("grid", &grid)

	win := component.CreateWindow(nil)

	grid.Text = win.Root().ObjectByName("text")

	win.Show()
	win.Wait()

	return nil
}
