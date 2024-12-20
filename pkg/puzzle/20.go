package puzzle

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/saaste/advent-of-code-2024/pkg/input"
	"github.com/x1m3/priorityQueue"
)

type Day20 struct{}

type PicoStepNode struct {
	Char     string
	Pos      Point2D
	Parent   *PicoStepNode
	Distance int64
}

func (p *PicoStepNode) HigherPriorityThan(p2 priorityQueue.Interface) bool {
	p2Node := p2.(*PicoStepNode)
	return p.Distance < p2Node.Distance
}

type PicoRaceGrid struct {
	Map   map[Point2D]*PicoStepNode
	Start *PicoStepNode
	End   *PicoStepNode
}

func (d Day20) Step1(puzzleInput string) string {
	// Steps without cheating: 9316
	grid := parsePicoRaceGrid(puzzleInput)

	pg := priorityQueue.New()
	pg.Push(grid.Start)
	lastNode := dijkstra(grid, pg)

	// Build the best route
	pathPoints := make(map[Point2D]*PicoStepNode)
	currentNode := lastNode
	for currentNode != nil {
		pathPoints[currentNode.Pos] = currentNode
		currentNode = currentNode.Parent
	}

	shortCuts := make(map[int64]int)
	totalShortCuts := 0

	// Search for short cuts
	for point, currentNode := range pathPoints {
		neighbors := []Point2D{
			{Y: -1}, // Top
			{X: 1},  // Right
			{Y: 1},  // Bottom
			{X: -1}, // Left
		}
		for _, vector := range neighbors {
			neighbor := point.Add(vector)

			neighborNode, found := grid.Map[neighbor]

			if !found || neighborNode.Char != "#" {
				// Neighbor is not a wall, no shortcut here
				continue
			}

			// Check if the cell behind the wall
			// is part of the later route
			behindCell := neighbor.Add(vector)
			if behindNode, found := pathPoints[behindCell]; found && behindNode.Distance > currentNode.Distance {
				// Yey! We found a shortcuts. Check how much we saved
				saved := behindNode.Distance - currentNode.Distance - 1 - 1
				if saved >= 100 {
					shortCuts[saved] += 1
					totalShortCuts++
				}

			}
		}
	}

	return fmt.Sprintf("%d", totalShortCuts)
}

func (d Day20) Step2(puzzleInput string) string {
	return ""
}

func dijkstra(grid *PicoRaceGrid, pg *priorityQueue.Queue) *PicoStepNode {
	handledNodes := make(map[Point2D]*PicoStepNode, 0)
	var currentNode *PicoStepNode

	for {
		currentNode = pg.Pop().(*PicoStepNode)

		neighbors := getNeighborNodes(grid, currentNode, true)
		for _, neighbor := range neighbors {
			// Skip handled neighbors
			if _, isHandled := handledNodes[neighbor.Pos]; isHandled {
				continue
			}

			// Update neighbor if this route is sorter then the previous value
			if currentNode.Distance+1 < neighbor.Distance {
				neighbor.Distance = currentNode.Distance + 1
				neighbor.Parent = currentNode
			}
			pg.Push(neighbor)
		}

		handledNodes[currentNode.Pos] = currentNode

		// We've reaced the end
		if currentNode.Pos.Equals(&grid.End.Pos) {
			break
		}
	}

	return currentNode
}

func getNeighborNodes(grid *PicoRaceGrid, node *PicoStepNode, dropWalls bool) []*PicoStepNode {
	neighborNodes := make([]*PicoStepNode, 0)
	neighbors := []Point2D{
		{Y: -1}, // Top
		{X: 1},  // Right
		{Y: 1},  // Bottom
		{X: -1}, // Left
	}

	for _, vector := range neighbors {
		neighborPos := node.Pos.Add(vector)
		if neighborNode, found := grid.Map[neighborPos]; found {
			if !dropWalls || neighborNode.Char != "#" {
				neighborNodes = append(neighborNodes, neighborNode)
			}
		}
	}
	return neighborNodes
}

func parsePicoRaceGrid(puzzleInput string) *PicoRaceGrid {
	grid := &PicoRaceGrid{
		Map: make(map[Point2D]*PicoStepNode),
	}

	lines := input.EachLineAsString(puzzleInput)
	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			point := Point2D{X: x, Y: y}
			switch char {
			case "#":
				fallthrough
			case ".":
				grid.Map[point] = &PicoStepNode{
					Char:     char,
					Pos:      point,
					Parent:   nil,
					Distance: math.MaxInt64,
				}
			case "S":
				grid.Map[point] = &PicoStepNode{
					Char:     ".",
					Pos:      point,
					Parent:   nil,
					Distance: 0,
				}
				grid.Start = grid.Map[point]
			case "E":
				grid.Map[point] = &PicoStepNode{
					Char:     ".",
					Pos:      point,
					Parent:   nil,
					Distance: math.MaxInt64,
				}
				grid.End = grid.Map[point]
			default:
				log.Fatalf("Unknown character %s", char)
			}

		}
	}

	return grid
}
