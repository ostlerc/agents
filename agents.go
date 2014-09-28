package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Worker interface {
	Work()         // Do a step of work
	Result() []int // Returns work done (inspected usually after all work is done)
	Clear()        // Reset state
}

type AntNode interface {
	Mark(int, int)
	Sniff() (int, int)
	Pos() (int, int)
	Neighbors() []AntNode
	Snatch() Food
	HasFood() bool
	Drop(Food)
	Enter()
	Exit()
}

func init() {
	rand.Seed(0)
}

type Food struct {
	life int
}

type Ant struct {
	id   int     // Ant id
	at   AntNode // Current node location
	food Food    // Holding food
}

func (a *Ant) Work() {
	a.at.Exit()
	a.at = a.Decide()
	a.at.Enter()
}

func (a *Ant) Decide() AntNode {
	p := make(map[int][]AntNode)
	x, y := a.at.Sniff()
	n := a.at.Neighbors()
	best := make([]int, 0)
	carryingFood := a.food.life-grid.Time > 0
	var choice AntNode

	//If We are not carrying food, check
	if !carryingFood {
		a.food = a.at.Snatch()
		carryingFood = a.food.life-grid.Time > 0
	}
	if carryingFood && a.at == grid.Nest {
		a.at.Drop(a.food)
		fmt.Println("dropped food at home ", grid.FoodQty, a.id)
		a.food = Food{0}
		carryingFood = false
	}

	for _, n := range n {
		if carryingFood {
			if n == grid.Nest {
				choice = n
				break
			}
		} else if n.HasFood() {
			choice = n
			break
		}
	}

	if choice == nil {
		for _, n := range n {
			m1, m2 := n.Sniff()

			if m2&a.id == a.id && !carryingFood { //lower priority somewhere we've already been
				m1 = 0
			}

			if _, ok := p[m1]; !ok {
				p[m1] = make([]AntNode, 0)
				best = append(best, m1)
			}
			p[m1] = append(p[m1], n)
		}
	}

	if carryingFood {
		y &^= a.id
		x = x + 1
	} else {
		if a.food.life != 0 { //food expired mid transfer so drop food and set trail to cold
			x = 1
			a.food = Food{0}
		}
		y = y | a.id
	}
	a.at.Mark(x, y)

	//fmt.Println("choice ", choice != nil, " carryingFood ", carryingFood, "x,y", x, y, " id ", a.id)

	if choice != nil {
		return choice
	}

	sort.Sort(sort.Reverse(sort.IntSlice(best)))
	v := p[best[0]]
	r := rand.Intn(len(v))
	return v[r]
}
