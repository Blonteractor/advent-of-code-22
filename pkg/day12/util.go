package day12

import (
	"container/heap"
	"math"
)

type Location struct {
	x int
	y int
}

type Node struct {
	height   rune
	gScore   float64
	hScore   float64
	location Location
	isEnd    bool
}

func (n Node) neighbors(grid [][]*Node) []*Node {
	var neighbors []*Node
	var node *Node
	x, y := n.location.x, n.location.y

	// Check up
	if y > 0 {
		node = grid[y-1][x]
		if (node.height - 1) <= n.height {
			neighbors = append(neighbors, node)
		}
	}

	// Check down
	if y < len(grid)-1 {
		node = grid[y+1][x]
		if (node.height - 1) <= n.height {
			neighbors = append(neighbors, node)
		}
	}

	// Check left
	if x > 0 {
		node = grid[y][x-1]
		if (node.height - 1) <= n.height {
			neighbors = append(neighbors, node)
		}
	}

	// Check right
	if x < len(grid[0])-1 {
		node = grid[y][x+1]
		if (node.height - 1) <= n.height {
			neighbors = append(neighbors, node)
		}
	}

	return neighbors
}

func sq(num float64) float64 {
	return math.Pow(num, 2)
}

func (n Node) heuristic(end Node) float64 {
	return math.Sqrt(sq(float64(n.location.x)-float64(end.location.x)) + sq(float64(n.location.y)-float64(end.location.y)))
}

func (n Node) fvalue() float64 {
	return n.gScore + n.hScore
}

func node(height rune, x int, y int) Node {
	return Node{
		height: height,
		gScore: math.Inf(1),
		hScore: math.Inf(1),
		location: Location{
			x: x,
			y: y,
		},
		isEnd: false,
	}
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].fvalue() < pq[j].fvalue()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

type Set map[*Node]struct{}

func (s *Set) Add(item *Node) bool {
	_, exists := (*s)[item]
	if exists {
		return false
	}
	(*s)[item] = struct{}{}
	return true
}

func aStar(start *Node, end *Node, grid [][]*Node) float64 {
	openSet := PriorityQueue{start}
	heap.Init(&openSet)

	start.gScore = 0
	start.hScore = start.heuristic(*end)

	visited := make(Set)
	visited.Add(start)

	for openSet.Len() != 0 {
		current := heap.Pop(&openSet).(*Node)

		// If the current is the end node, we found the path, return the gScore of end node (number of steps)
		if current.isEnd {
			return current.gScore
		}

		for _, neighbor := range current.neighbors(grid) {
			tentative_gScore := current.gScore + 1
			if tentative_gScore < neighbor.gScore {
				neighbor.gScore = tentative_gScore
				neighbor.hScore = neighbor.heuristic(*end)
	
				// Add neighbor in open list if its not visited
				if visited.Add(neighbor) {
					heap.Push(&openSet, neighbor)
				}
			}
		}
	}
	return math.Inf(1)
}
