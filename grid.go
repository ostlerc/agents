package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"time"

	"gopkg.in/qml.v1"
)

type Grid struct {
	Rows         qml.Object
	Cols         qml.Object
	Grid         qml.Object
	StatusText   qml.Object
	DefFoodCnt   qml.Object
	FoodTime     qml.Object
	FoodCnt      qml.Object
	LifeCnt      qml.Object
	RunBtn       qml.Object
	StepBtn      qml.Object
	PauseBtn     qml.Object
	SimText      qml.Object
	FoodText     qml.Object
	DelaySpinner qml.Object
	GenerateFood qml.Object

	TileComp *Tile
	Nest     *Tile
	Selected *Tile

	Tiles []*Tile

	Edited   bool
	ColCount int
	RowCount int
	FoodQty  int
	MaxFood  int

	StopChan  chan bool
	PauseChan chan bool

	Time int

	Ants []*Ant
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
	g.Nest = g.Tiles[i]
	if g.RunBtn != nil {
		g.RunBtn.Set("enabled", g.Nest != nil)
	}
}

func (g *Grid) ClearNest() {
	g.Nest = nil
	if g.RunBtn != nil {
		g.RunBtn.Set("enabled", false)
	}
}

func (g *Grid) ResetStatus() {
	g.FoodCnt.Set("visible", false)
	g.LifeCnt.Set("visible", false)
	g.StatusText.Set("text", "Click the grid cells to make a Nest, food, and walls. Right click to inspect.")
}

func (g *Grid) SetSelected(i int) {
	if g.Selected != nil {
		g.Selected.Object.Set("selected", false)
		if g.Tiles[i] == g.Selected {
			g.Selected = nil
			g.ResetStatus()
			return
		}
	}
	g.Selected = g.Tiles[i]
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
	return fmt.Sprintf("%v %v %v", name, t.x, t.y)
}

func (g *Grid) createTile() *Tile {
	tile := &Tile{
		Object:   g.TileComp.Object.Create(nil),
		diagonal: true,
		x:        1,
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
	if _, err := os.Stat(filename); err != nil {
		filename = filename[7:]
	}
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

func (g *Grid) Assign(name string, o qml.Object) {
	switch name {
	case "Run":
		g.RunBtn = o
	case "Pause":
		g.PauseBtn = o
	case "Step":
		g.StepBtn = o
	case "simStatus":
		g.SimText = o
	case "foodStatus":
		g.FoodText = o
	case "delaySpinner":
		g.DelaySpinner = o
	case "foodGen":
		g.GenerateFood = o
	default:
		panic("Invalid name " + name)
	}
}

func (g *Grid) RunClicked() {
	if g.RunBtn.String("text") == "Stop" {
		go func() {
			g.StopChan <- true
		}()
		return
	}
	g.RunBtn.Set("text", "Stop")
	g.PauseBtn.Set("enabled", true)

	g.Ants = make([]*Ant, 0)
	for i := 1; i < g.Nest.Object.Int("count")+1; i++ {
		a := &Ant{
			id: int(math.Pow(2, float64(i))),
			at: g.Nest,
		}
		fmt.Println("Created ant with id: ", a.id)
		g.Ants = append(g.Ants, a)
		g.Nest.Enter(false)
	}

	g.MaxFood = 0
	for _, t := range g.Tiles {
		g.MaxFood += t.Food()
		t.Mark(0, 0)
	}

	fmt.Println("Gathering", g.MaxFood, "Foods")

	g.Time = 0
	g.FoodQty = 0
	g.StopChan = make(chan bool)
	g.PauseChan = make(chan bool)
	go func() {
		for {
			select {
			case <-g.StopChan:
				g.PauseBtn.Set("enabled", false)
				g.RunBtn.Set("text", "Run")
				g.PauseBtn.Set("text", "Pause")
				g.ClearGrid()
				return
			case <-g.PauseChan:
				v := g.PauseBtn.String("text")
				if v == "Pause" {
					g.PauseBtn.Set("text", "Unpause")
					g.StepBtn.Set("enabled", true)
				} else {
					g.PauseBtn.Set("text", "Pause")
					g.StepBtn.Set("enabled", false)
				}
				select { //pause means no default in select
				case <-g.PauseChan:
					v := g.PauseBtn.String("text")
					if v == "Pause" {
						g.PauseBtn.Set("text", "Unpause")
						g.StepBtn.Set("enabled", true)
					} else {
						g.PauseBtn.Set("text", "Pause")
						g.StepBtn.Set("enabled", false)
					}
				case <-g.StopChan:
					g.RunBtn.Set("text", "Run")
					g.PauseBtn.Set("text", "Pause")
					g.ClearGrid()
					return
				}
			default:
				g.Step()
			}
		}
	}()
}

func (g *Grid) Step() {
	genFood := g.GenerateFood.Bool("checked")
	for _, tile := range g.Tiles {
		if t := tile.Object.Int("type"); t == 3 {
			l := tile.Object.Int("life")
			c := tile.Object.Int("count")
			if l-g.Time <= 0 || c == 0 {
				g.MaxFood -= c
				tile.Object.Set("count", 0)
				tile.Object.Set("type", 0)
				tile.Object.Set("life", 0)
				fmt.Println(c, "Food Expired", g.Time)
			}
		} else if genFood && t == 0 && rand.Intn(5000) == 5 {
			//randomly generate food
			tile.Object.Set("type", 3)
			c := rand.Intn(15)
			g.MaxFood += c
			tile.Object.Set("count", c)
			tile.Object.Set("life", rand.Intn(100)+g.Time)
		}
	}
	for _, ant := range g.Ants {
		ant.Work()
	}

	g.SimText.Set("text", g.Time)
	g.FoodText.Set("text", g.FoodQty)
	time.Sleep(time.Duration(g.DelaySpinner.Int("value")) * time.Millisecond)
	g.Time++

	if g.FoodQty == g.MaxFood && !g.GenerateFood.Bool("checked") {
		fmt.Println("Gathered", g.FoodQty, "in", g.Time, "steps")
		go func() {
			g.StopChan <- true
		}()
		time.Sleep(time.Millisecond) //give some time to process the stop
	}
}

func (g *Grid) PauseClicked() {
	go func() {
		g.PauseChan <- true
	}()
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
	g.ResetStatus()
	g.Selected = nil

	fmt.Println("Building a", g.RowCount, g.ColCount, "grid")
	size := g.RowCount * g.ColCount
	g.Tiles = make([]*Tile, size, size)
	for n := 0; n < size; n++ {
		g.Tiles[n] = g.createTile()
		g.Tiles[n].Object.Set("index", n)
	}
}

func (g *Grid) ClearGrid() {
	g.Edited = true
	for _, v := range g.Tiles {
		v.Object.Set("solution", false)
		t := v.Object.Int("type")
		c := 0
		l := 0
		if t == 2 {
			c = 10
		} else if t == 3 {
			c = g.DefFoodCnt.Int("value")
			l = g.FoodTime.Int("value")
		}
		v.Object.Set("count", c)
		v.Object.Set("life", l)
		v.Object.Set("pcount", 0)
		v.Object.Set("selected", 0)
		v.Object.Set("antcount", 0)
	}
}
