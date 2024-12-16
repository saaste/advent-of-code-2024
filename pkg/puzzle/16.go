package puzzle

import (
	"container/heap"
	"fmt"
	"math"
	"strings"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day16 struct{}

type MazeGraph struct {
	NodeMap  map[string]*Node
	NodeList []*Node
	Start    Point2D
	End      Point2D
	Size     int
}

type Node struct {
	Pos    Point2D
	Top    *Node
	Right  *Node
	Bottom *Node
	Left   *Node
	Facing int

	Priority int
	Index    int
	Cost     int
	Parent   *Node
}

func (d Day16) Step1(puzzleInput string) string {
	mazeGraph := parseMazeGraph(puzzleInput)

	path, _ := aStar(mazeGraph)

	points := 0
	var previousNode *Node
	for _, node := range path {
		if previousNode != nil {
			points += previousNode.Pos.Distance(&node.Pos)
			if node.Facing != previousNode.Facing {
				points += 1000
			}
		}
		previousNode = node
	}

	return fmt.Sprintf("%d", points)
}

func (d Day16) Step2(puzzleInput string) string {
	return ""
}

func heuristic(a, b Point2D, _ int) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}

const (
	FacingRight = iota
	FacingBottom
	FacingLeft
	FacingTop
)

func aStar(graph *MazeGraph) ([]*Node, int) {

	nodesInHandling := &PriorityQueue{}
	heap.Init(nodesInHandling)

	startNode := graph.NodeMap[graph.Start.ToKey()]
	startNode.Facing = FacingRight

	heap.Push(nodesInHandling, startNode)

	visitedNodes := make(map[string]bool)

	for nodesInHandling.Len() > 0 {
		// Pick the node with the lowest score
		currentNode := heap.Pop(nodesInHandling).(*Node)

		// We've reached the end! Wohoo!
		if currentNode.Pos.Equals(&graph.End) {
			path := make([]*Node, 0)
			for node := currentNode; node != nil; node = node.Parent {
				path = append([]*Node{node}, path...)
			}
			return path, currentNode.Cost
		}

		// Mark current node visited
		visitedNodes[currentNode.Pos.ToKey()] = true

		// Check neighbors
		handleNeighbord(graph, nodesInHandling, visitedNodes, currentNode, currentNode.Top, FacingTop)
		handleNeighbord(graph, nodesInHandling, visitedNodes, currentNode, currentNode.Right, FacingRight)
		handleNeighbord(graph, nodesInHandling, visitedNodes, currentNode, currentNode.Bottom, FacingBottom)
		handleNeighbord(graph, nodesInHandling, visitedNodes, currentNode, currentNode.Left, FacingLeft)

	}
	// No route :(
	return nil, -1

}

func handleNeighbord(graph *MazeGraph, nodesInHandling *PriorityQueue, visitedNodes map[string]bool, currentNode *Node, neighbor *Node, facing int) {
	if neighbor != nil && !visitedNodes[neighbor.Pos.ToKey()] {
		turnCost := 0
		if currentNode.Facing != facing {
			turnCost = 1000
		}

		currentCost := currentNode.Cost
		distanceFromPrevious := currentNode.Pos.Distance(&neighbor.Pos)

		newCost := currentCost + distanceFromPrevious + turnCost
		priority := newCost + heuristic(neighbor.Pos, graph.End, turnCost)

		// Check if neighbor is already in the handling queue
		isInHandling := false
		for _, node := range *nodesInHandling {
			if node.Pos.Equals(&neighbor.Pos) {
				isInHandling = true
				if newCost < node.Cost {
					node.Facing = facing
					nodesInHandling.Update(node, newCost, priority)
				}
				break
			}
		}

		if !isInHandling {
			neighbor.Cost = newCost
			neighbor.Priority = priority
			neighbor.Parent = currentNode
			neighbor.Facing = facing
			heap.Push(nodesInHandling, neighbor)
		}
	}
}

func parseMazeGraph(puzzleInput string) *MazeGraph {
	mazeGraph := &MazeGraph{
		NodeMap:  make(map[string]*Node),
		NodeList: make([]*Node, 0),
	}

	// First make simple 2D grid as a map
	grid := make(map[string]string, 0)
	lines := input.EachLineAsString(puzzleInput)
	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			point := Point2D{X: x, Y: y}
			if char == "S" {
				mazeGraph.Start = point
				grid[point.ToKey()] = "."
			} else if char == "E" {
				mazeGraph.End = point
				grid[point.ToKey()] = "."
			} else {
				grid[point.ToKey()] = char
			}
		}
	}

	var prevChar string
	var prevLeftNode *Node

	// Now we can check each cell and check if it is a node
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines); x++ {
			pos := Point2D{X: x, Y: y}
			char := grid[pos.ToKey()]

			top := pos.Add(Point2D{Y: -1})
			bottom := pos.Add(Point2D{Y: 1})
			right := pos.Add(Point2D{X: 1})

			if char == "#" {
				// Ignore walls
				prevChar = char
				prevLeftNode = nil
				continue
			}

			// If previous cell was a wall and this one isn't
			if prevChar == "#" && char == "." {
				// If there is a wall above or below, this is a node
				if grid[top.ToKey()] == "#" || grid[bottom.ToKey()] == "#" {
					node := &Node{Pos: pos}

					if grid[top.ToKey()] == "." {
						topNode := findNextNodeAbove(mazeGraph, node)
						topNode.Bottom = node
						node.Top = topNode
					}

					mazeGraph.NodeList = append(mazeGraph.NodeList, node)
					mazeGraph.NodeMap[pos.ToKey()] = node
					prevChar = char
					prevLeftNode = node
					continue
				}

				// If right side is open, this is a start of a horizontal corridor
				if grid[right.ToKey()] == "." {
					node := &Node{Pos: pos}

					if grid[top.ToKey()] == "." {
						topNode := findNextNodeAbove(mazeGraph, node)
						topNode.Bottom = node
						node.Top = topNode
					}

					mazeGraph.NodeList = append(mazeGraph.NodeList, node)
					mazeGraph.NodeMap[pos.ToKey()] = node
					prevChar = char
					prevLeftNode = node
					continue
				}

				// Otherwise this is vertical corridor
				prevChar = char
				prevLeftNode = nil
				continue
			}

			if prevChar == "." && char == "." {

				// If next cell is a wall, this is the end of the corridor
				// so it must be a node
				if grid[right.ToKey()] == "#" {
					node := &Node{Pos: pos, Left: prevLeftNode}
					if prevLeftNode != nil {
						prevLeftNode.Right = node
						node.Left = prevLeftNode
					}

					// If cell above is open, connect to the node above
					if grid[top.ToKey()] == "." {
						nodeAbove := findNextNodeAbove(mazeGraph, node)
						node.Top = nodeAbove
						nodeAbove.Bottom = node

					}

					mazeGraph.NodeList = append(mazeGraph.NodeList, node)
					mazeGraph.NodeMap[pos.ToKey()] = node
					prevChar = char
					prevLeftNode = nil
					continue
				}

				// If previous character was a free space and this one is too, check
				// if we can go up or down. If yes, this is a node.
				if grid[top.ToKey()] == "." || grid[bottom.ToKey()] == "." {
					node := &Node{Pos: pos, Left: prevLeftNode}

					if prevLeftNode != nil {
						prevLeftNode.Right = node
						node.Left = prevLeftNode
					}

					// If cell above is open, connect to the node above
					if grid[top.ToKey()] == "." {
						nodeAbove := findNextNodeAbove(mazeGraph, node)
						node.Top = nodeAbove
						nodeAbove.Bottom = node
					}

					mazeGraph.NodeList = append(mazeGraph.NodeList, node)
					mazeGraph.NodeMap[pos.ToKey()] = node
					prevChar = char
					prevLeftNode = node
					continue
				}

				// This is not a node
				prevChar = char
			}

		}
	}
	mazeGraph.Size = len(lines)
	return mazeGraph
}

func findNextNodeAbove(mazeGraph *MazeGraph, currentNode *Node) *Node {
	currentPos := currentNode.Pos
	// Move up until we find a node
	for y := currentNode.Pos.Y; y > 0; y-- {
		currentPos = currentPos.Add(Point2D{Y: -1})

		for _, node := range mazeGraph.NodeList {
			if node.Pos.Equals(&currentPos) {
				return node

			}
		}
	}
	return nil
}

// -----------------
// Priority queue
// --------------------
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(a, b int) bool {
	return pq[a].Priority < pq[b].Priority
}

func (pq PriorityQueue) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]
	pq[a].Index = a
	pq[b].Index = b
}

func (pq *PriorityQueue) Push(n interface{}) {
	node := n.(*Node)
	node.Index = len(*pq)
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)

	node := old[n-1]
	old[n-1] = nil

	node.Index = -1
	*pq = old[0 : n-1]
	return node
}

func (pq *PriorityQueue) Update(node *Node, cost, priority int) {
	node.Cost = cost
	node.Priority = priority
	heap.Fix(pq, node.Index)
}
