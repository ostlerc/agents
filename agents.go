package main

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
	Drop(Food)
	Enter()
	Exit()
}

type Food struct {
	life int
}

type Ant struct {
	id   int     // Ant id
	at   AntNode // Current node location
	food *Food   // Holding food
}
