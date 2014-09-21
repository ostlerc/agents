package main

type Graph interface {
	CalculatePath(start, end Node) ([]Node, error)
	WorkDone() int
}

type Node interface {
	Neighbors() []Node
	Dist(neighbor Node) float64
	EstimatedCost(goal Node) float64
}

func NewAstar(nodes []Node) Graph {
	return &astar{Nodes: nodes}
}

type astar struct {
	Nodes []Node
	work  int
}

func (a *astar) WorkDone() int {
	return a.work
}

func (a *astar) CalculatePath(start, goal Node) ([]Node, error) {
	a.work = 0
	closedset := make(map[Node]bool)      // The set of nodes already evaluated.
	openset := map[Node]bool{start: true} // The set of tentative nodes to be evaluated, initially containing the start node
	came_from := make(map[Node]Node)      // The map of navigated nodes.

	g_score := make(map[Node]float64)
	f_score := make(map[Node]float64)
	g_score[start] = 0 // Cost from start along best known path.
	// Estimated total cost from start to goal through y.
	f_score[start] = start.EstimatedCost(goal)

	for len(openset) > 0 {
		current := lowest(f_score)
		a.work++
		if current == goal {
			return reconstructPath(came_from, goal), nil
		}
		delete(f_score, current)
		delete(openset, current)
		closedset[current] = true

		for _, neighbor := range current.Neighbors() {
			if _, ok := closedset[neighbor]; ok {
				continue
			}
			tentative_g_score := g_score[current] + current.Dist(neighbor)

			if _, ok := openset[neighbor]; !ok || tentative_g_score < g_score[neighbor] {
				came_from[neighbor] = current
				g_score[neighbor] = tentative_g_score
				f_score[neighbor] = g_score[neighbor] + neighbor.EstimatedCost(goal)
				if _, ok := openset[neighbor]; !ok {
					openset[neighbor] = true
				}
			}
		}

	}
	return []Node{}, nil
}

func reconstructPath(came_from map[Node]Node, goal Node) []Node {
	res := make([]Node, 0, 500)
	res = append(res, goal)
	at := goal
	for {
		v, ok := came_from[at]
		if !ok {
			break
		}
		res = append(res, v)
		at = v

	}
	return res
}

func lowest(nodes map[Node]float64) Node {
	min := float64(-1)
	var res Node
	for k, v := range nodes {
		if min == -1 || v < min {
			min = v
			res = k
		}
	}
	return res
}
