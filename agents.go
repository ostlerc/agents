package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Worker interface {
	Work() // Do a step of work
}

type AntNode interface {
	Mark(int, int)
	Sniff() (int, int)
	Pos() (int, int)
	Neighbors() []AntNode
	Snatch() Food
	Food() int
	Drop(Food)
	Enter(bool)
	Exit()
}

func init() {
	rand.Seed(52 * 1332)
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
	a.at.Enter(a.food.life-grid.Time > 0)
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
		if a.food.life > 0 {
			fmt.Println("My food expired!")
			grid.MaxFood--
		}
		a.food = a.at.Snatch()
		carryingFood = a.food.life-grid.Time > 0
		if !carryingFood && a.food.life > 0 {
			fmt.Println("My food expired!")
			grid.MaxFood--
		}
	}
	if carryingFood && a.at == grid.Nest {
		a.at.Drop(a.food)
		fmt.Println("food drop", grid.FoodQty, grid.Time)
		a.food = Food{0}
		carryingFood = false
	}

	for _, n := range n {
		_, m2 := n.Sniff()

		if carryingFood {
			if n == grid.Nest {
				choice = n
				break
			}
			if m2&a.id == a.id {
				if choice == nil {
					choice = n
				}
			}
		} else if n.Food() > 0 {
			choice = n
		}
	}

	if choice == nil && carryingFood { //food trail gone cold
		x = 1
		y &^= a.id
		carryingFood = false
	}

	if choice == nil {
		for _, n := range n {
			m1, m2 := n.Sniff()

			if m2&a.id == a.id { //lower priority somewhere we've already been
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
		y |= 1 //mark as food trail, remove id token
		y &^= a.id
		x = x + 1
	} else {
		if a.food.life != 0 { //food expired mid transfer so drop food and set trail to cold
			y &^= a.id
		} else {
			y = y | a.id
		}
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
