package puzzle

import (
	"fmt"
	"math/bits"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day12 struct{}

func (d Day12) Step1(puzzleInput string) string {
	mapSize := len(input.EachLineAsString(puzzleInput))
	field := input.CharacterMap(puzzleInput)

	foundAreas := make([]*Area, 0)
	visitedPlots := make(map[string]bool)
	for y := 0; y < mapSize; y++ {
		for x := 0; x < mapSize; x++ {
			currentPoint := Point2D{X: x, Y: y}
			foundAreas, _ = findAreas(field, currentPoint, nil, visitedPlots, foundAreas)
		}
	}

	totalPrice := 0
	for _, area := range foundAreas {
		totalPrice += area.Area * area.Perimeter
	}

	return fmt.Sprintf("%d", totalPrice)
}

func (d Day12) Step2(puzzleInput string) string {
	mapSize := len(input.EachLineAsString(puzzleInput))
	field := input.CharacterMap(puzzleInput)

	foundAreas := make([]*Area, 0)
	visitedPlots := make(map[string]bool)
	for y := 0; y < mapSize; y++ {
		for x := 0; x < mapSize; x++ {
			currentPoint := Point2D{X: x, Y: y}
			foundAreas, _ = findAreas(field, currentPoint, nil, visitedPlots, foundAreas)
		}
	}

	totalPrice := 0

	for _, area := range foundAreas {
		totalPrice += area.Area * calculateSides(area)
	}

	return fmt.Sprintf("%d", totalPrice)
}

type Area struct {
	Code      string
	Plots     map[string]*Plot
	PlotSlice []*Plot
	Area      int
	Perimeter int
}

const (
	Top    = byte(0b00000001)
	Right  = byte(0b00000010)
	Bottom = byte(0b00000100)
	Left   = byte(0b00001000)
)

type Plot struct {
	point       Point2D
	IsPerimeter bool
	Walls       byte
	CheckNumber int
}

func (p *Plot) WallCount() int {
	return bits.OnesCount(uint(p.Walls))
}

func (p *Plot) NewWallCount(other Plot) int {
	difference := p.Walls ^ other.Walls
	newWallsInOther := other.Walls & difference
	return bits.OnesCount(uint(newWallsInOther))
}

func (p *Plot) SharedWallCount(other Plot) int {
	return bits.OnesCount(uint(p.Walls & other.Walls))
}

func findAreas(field map[string]string, currentPlot Point2D, currentArea *Area, visitedPlots map[string]bool, foundAreas []*Area) ([]*Area, bool) {
	currentPlotKey := currentPlot.ToKey()

	// Skip plots that are not in the map. However,
	// we need to increase the perimeter by one
	if _, found := field[currentPlotKey]; !found {
		if currentArea != nil {
			currentArea.Perimeter++
		}
		return foundAreas, false
	}

	currentPlotCode := field[currentPlotKey]

	// Skip plots that are already checked, but increase
	// the perimeter if plot does not belong to the area
	if _, found := visitedPlots[currentPlot.ToKey()]; found {
		belongsToTheArea := true
		if currentArea != nil && currentPlotCode != currentArea.Code {
			currentArea.Perimeter++
			belongsToTheArea = false
		}
		return foundAreas, belongsToTheArea
	}

	if currentArea == nil {
		// Create a new area if needed
		currentArea = &Area{
			Code:      currentPlotCode,
			Plots:     make(map[string]*Plot),
			PlotSlice: make([]*Plot, 0),
			Area:      0,
			Perimeter: 0,
		}
		foundAreas = append(foundAreas, currentArea)
	}

	// This plot does not belong to the existing area so
	// we can just skip this for now. However, we need to
	// increase the perimeter by one
	if currentArea.Code != currentPlotCode {
		currentArea.Perimeter++
		return foundAreas, false
	}

	// Otherwise we can add the plot to the existing area,
	// increase the area size by one and mark the plot as
	// visited
	visitedPlots[currentPlotKey] = true
	currentArea.Area++

	// Check neighboring plots
	neighbors := []Point2D{
		{X: currentPlot.X, Y: currentPlot.Y - 1}, // Top
		{X: currentPlot.X, Y: currentPlot.Y + 1}, // Bottom
		{X: currentPlot.X + 1, Y: currentPlot.Y}, // Right
		{X: currentPlot.X - 1, Y: currentPlot.Y}, // Left
	}

	// Store information about plot's perimeter and walls
	isPerimeter := false
	walls := byte(0b00000000)

	for i, neighbor := range neighbors {
		if _, belongsToArea := findAreas(field, neighbor, currentArea, visitedPlots, foundAreas); !belongsToArea {
			isPerimeter = true
			if i == 0 {
				walls = walls | Top
			} else if i == 1 {
				walls = walls | Bottom
			} else if i == 2 {
				walls = walls | Right
			} else {
				walls = walls | Left
			}
		}
	}

	// Add
	if isPerimeter {
		plot := Plot{point: Point2D{X: currentPlot.X, Y: currentPlot.Y}, IsPerimeter: isPerimeter, Walls: walls}
		currentArea.Plots[currentPlotKey] = &plot
		currentArea.PlotSlice = append(currentArea.PlotSlice, &plot)
	}

	// We are done with this this plot
	return foundAreas, true
}

func calculateSides(area *Area) int {
	sideCount := 0
	checkedPlots := make(map[string]bool)
	checkNumber := 0

	// Simple shapes go through in one go, but with some shapes
	// there can be isolated plots that are not part of the perimeter
	// walk. This is why we loop through all plots just to make sure
	// we've handled everything
	for _, currentPlot := range area.PlotSlice {
		if _, found := checkedPlots[currentPlot.point.ToKey()]; found {
			continue
		}

		sideCount += walkPerimeter(area, currentPlot, nil, checkNumber, checkedPlots, -1)
	}
	return sideCount
}

// Walk the perimeter and count sides
func walkPerimeter(area *Area, currentPlot *Plot, previousPlot *Plot, checkNumber int, checkedPlots map[string]bool, sideCount int) int {

	if sideCount == -1 {
		// Initialize the walk with sides of the first plot
		sideCount = currentPlot.WallCount()
	} else if previousPlot != nil {
		// Otherwise add new walls compared to the previos plot
		sideCount += previousPlot.NewWallCount(*currentPlot)
	}

	// Store to cache
	checkedPlots[currentPlot.point.ToKey()] = true

	// Set the check number
	currentPlot.CheckNumber = checkNumber

	// Loop through neighbors
	neighbors := getNeighbors(currentPlot)
	for _, neighborPoint := range neighbors {
		neighborKey := neighborPoint.ToKey()

		// Don't check the previous plot
		if previousPlot != nil && neighborPoint == previousPlot.point {
			continue
		}

		// Neighbor is already checked, but...
		if _, checked := checkedPlots[neighborKey]; checked {
			if neighborPlot, found := area.Plots[neighborPoint.ToKey()]; found && neighborPlot.CheckNumber < currentPlot.CheckNumber {
				// ...subtract common sides so that we don't count them twice
				sideCount -= neighborPlot.SharedWallCount(*currentPlot)
			}
			// ...and then we can move to the next neighbour
			continue
		}

		if neighborPlot, found := area.Plots[neighborPoint.ToKey()]; found {
			// ...recursioooooooon with all neighbors....
			sideCount = walkPerimeter(area, neighborPlot, currentPlot, checkNumber+1, checkedPlots, sideCount)
		}
	}

	return sideCount
}

func getNeighbors(plot *Plot) []Point2D {
	return []Point2D{
		{X: plot.point.X, Y: plot.point.Y - 1}, // Top
		{X: plot.point.X + 1, Y: plot.point.Y}, // Right
		{X: plot.point.X, Y: plot.point.Y + 1}, // Bottom
		{X: plot.point.X - 1, Y: plot.point.Y}, // Left
	}
}
