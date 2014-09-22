package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"gopkg.in/qml.v1"
)

type Grid struct {
	Rows       qml.Object
	Cols       qml.Object
	Grid       qml.Object
	StatusText qml.Object
	DefFoodCnt qml.Object
	FoodTime   qml.Object
	RunBtn     qml.Object
	FoodCnt    qml.Object
	LifeCnt    qml.Object

	TileComp *Tile
	Nest     *Tile
	Selected *Tile

	Tiles []Tile

	Edited   bool
	ColCount int
	RowCount int
}

type JSONGrid struct {
	Rows  int        `json:"rows"`
	Cols  int        `json:"cols"`
	Food  int        `json:"food"`
	Life  int        `json:"life"`
	Tiles []JSONTile `json:"tiles"`
}

func (g *Grid) SetCount(i int) {
	if g.Selected == nil {
		fmt.Println("Error setting on nil selected (SetCount)")
		return
	}

	g.Selected.Object.Set("count", i)
}

func (g *Grid) SetLife(i int) {
	if g.Selected == nil {
		fmt.Println("Error setting on nil selected (SetLife)")
		return
	}

	g.Selected.Object.Set("life", i)
}

func (g *Grid) FoodCount() int {
	return g.DefFoodCnt.Int("value")
}

func (g *Grid) FoodLife() int {
	return g.FoodTime.Int("value")
}

func (g *Grid) SetNest(i int) {
	g.Nest = &g.Tiles[i]
	g.RunBtn.Set("visible", g.Nest != nil)
}

func (g *Grid) ClearNest() {
	g.Nest = nil
	g.RunBtn.Set("visible", false)
}

func (g *Grid) ResetStatus() {
	g.FoodCnt.Set("visible", false)
	g.LifeCnt.Set("visible", false)
	g.StatusText.Set("text", "Click the grid cells to make a Nest, food, and walls. Right click to inspect.")
}

func (g *Grid) SetSelected(i int) {
	if g.Selected != nil {
		g.Selected.Object.Set("selected", false)
		if &g.Tiles[i] == g.Selected {
			g.Selected = nil
			g.ResetStatus()
			return
		}
	}
	g.Selected = &g.Tiles[i]
	g.Selected.Object.Set("selected", true)
	g.UpdateStatus()
}

func (g *Grid) UpdateStatus() {
	if g.Selected == nil {
		return
	}
	g.FoodCnt.Set("visible", false)
	g.LifeCnt.Set("visible", false)
	t := g.Selected.Object.Int("type")
	g.SetStatus(g.StatusFromTile(g.Selected))
	switch t {
	case 0: //open
	case 1: //wall
	case 3: //food
		g.LifeCnt.Set("visible", true)
		g.LifeCnt.Set("value", g.Selected.Object.Int("life"))
		fallthrough
	case 2: //nest
		g.FoodCnt.Set("visible", true)
		g.FoodCnt.Set("value", g.Selected.Object.Int("count"))
	}
}

func (g *Grid) SetStatus(s interface{}) {
	g.StatusText.Set("text", s)
}

func (g *Grid) StatusFromTile(t *Tile) string {
	name := "open"
	switch t.Object.Int("type") {
	case 1:
		name = "wall"
	case 2:
		name = "nest"
	case 3:
		name = "food"
	}
	return fmt.Sprintf("%v", name)
}

func (g *Grid) createTile() Tile {
	tile := Tile{
		Object:   g.TileComp.Object.Create(nil),
		diagonal: true,
	}
	tile.Object.Set("parent", g.Grid)
	return tile
}

func (g *Grid) SaveGrid(filename string) {
	filename = filename[7:]
	jg := &JSONGrid{
		Rows: g.RowCount,
		Cols: g.ColCount,
		Food: g.FoodCount(),
		Life: g.FoodLife(),
	}
	tiles := make([]JSONTile, 0, jg.Rows*jg.Cols)

	for _, v := range g.Tiles {
		ttype := v.Object.Int("type")
		if ttype == 0 { //skip open nodes
			continue
		}
		t := JSONTile{
			Type:  ttype,
			Count: v.Object.Int("count"),
			Life:  v.Object.Int("life"),
			Index: v.Object.Int("index"),
		}
		tiles = append(tiles, t)
	}

	jg.Tiles = tiles

	dat, err := json.Marshal(jg)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(dat))
	err = ioutil.WriteFile(filename, dat, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully Saved", filename)
}

func (g *Grid) LoadGrid(filename string) {
	filename = filename[7:]
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	var newg JSONGrid
	err = json.Unmarshal(dat, &newg)
	if err != nil {
		fmt.Println(err)
		return
	}
	g.Rows.Set("value", newg.Rows)
	g.Cols.Set("value", newg.Cols)
	g.DefFoodCnt.Set("value", newg.Food)
	g.FoodTime.Set("value", newg.Life)
	g.BuildGrid()
	for _, t := range newg.Tiles {
		g.Tiles[t.Index].Object.Set("type", t.Type)
		g.Tiles[t.Index].Object.Set("life", t.Life)
		g.Tiles[t.Index].Object.Set("count", t.Count)

		switch t.Type {
		case 2: //next
			g.SetNest(t.Index)
		}
	}
	fmt.Println("Successfully Loaded", filename)
}

func (g *Grid) BuildGrid() {
	g.Edited = true
	for _, b := range g.Tiles {
		b.Object.Set("visible", false)
		b.Object.Destroy()
	}
	g.ClearNest()
	g.RowCount = g.Rows.Int("value")
	g.ColCount = g.Cols.Int("value")
	g.Grid.Set("columns", g.ColCount)
	g.RunBtn.Set("visible", false)
	g.ResetStatus()
	g.Selected = nil

	fmt.Println("Building a", g.RowCount, g.ColCount, "grid")
	size := g.RowCount * g.ColCount
	g.Tiles = make([]Tile, size, size)
	for n := 0; n < size; n++ {
		g.Tiles[n] = g.createTile()
		g.Tiles[n].Object.Set("index", n)
	}
}

func (g *Grid) ClearGrid() {
	g.Edited = true
	for _, v := range g.Tiles {
		v.Object.Set("solution", false)
	}
}

func (g *Grid) colorSolution(objs []AntNode) {
	for _, v := range objs {
		v.(*Tile).Object.Set("solution", true)
	}
}
