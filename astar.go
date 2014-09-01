package main

type Graph interface {
	CalculatePath(start, end Node) ([]Node, error)
}

type Node interface {
	Edges() []Node
	Cost(goal Node) float64
	Dist(goal Node) float64
}

type astar struct{}

func (a *astar) CalculatePath(start, end Node) ([]Node, error) {
	return []Node{start, end}, nil
}

func New() Graph {
	return &astar{}
}
