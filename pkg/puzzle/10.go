package puzzle

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day10 struct{}

func (d Day10) Step1(puzzleInput string) string {
	hikingMap := parseHikingMap(puzzleInput)

	// Calculate the score of each trail head
	mapScore := 0
	for _, point := range hikingMap.StartingPoints {
		mapScore += calculateRouteScore(hikingMap, point, map[string]Point2D{}, 0, true)
	}

	return fmt.Sprintf("%d", mapScore)
}

func (d Day10) Step2(puzzleInput string) string {
	hikingMap := parseHikingMap(puzzleInput)

	// Calculate the score of each trail head
	mapScore := 0
	for _, point := range hikingMap.StartingPoints {
		// mapScore += calculateUniqueRouteScore(hikingMap, point, 0)
		mapScore += calculateRouteScore(hikingMap, point, map[string]Point2D{}, 0, false)
	}

	return fmt.Sprintf("%d", mapScore)
}

type HikingMap struct {
	Map            map[string]int
	StartingPoints []Point2D
}

func calculateRouteScore(hikingMap HikingMap, currentLocation Point2D, endings map[string]Point2D, currentScore int, checkUniqueRoutes bool) int {
	// Get the height of the current location
	currentLocationHeight, found := hikingMap.Map[currentLocation.ToKey()]
	if !found {
		log.Fatalf("Invalid direction. How on earth we ended up here?! - %+v", currentLocation)
	}

	// Check if found an ending
	if currentLocationHeight == 9 {
		if checkUniqueRoutes {
			// Count only unique endings (Step 1)
			if _, endingAlreadyFound := endings[currentLocation.ToKey()]; !endingAlreadyFound {
				// We have a new ending! Add it to the list and return an updated score
				endings[currentLocation.ToKey()] = currentLocation
				return currentScore + 1
			} else {
				// We've already reached this ending. Return current score
				return currentScore
			}
		} else {
			// Count all possible routes (Step 2)
			return currentScore + 1
		}
	}

	// Find directions we can go
	possibleDirections := findPossibleDirections(hikingMap, currentLocation)

	// Start traversing each direction
	for _, nextPoint := range possibleDirections {
		currentScore = calculateRouteScore(hikingMap, nextPoint, endings, currentScore, checkUniqueRoutes)
	}

	// All directions checked, return current score
	return currentScore
}

func findPossibleDirections(hikingMap HikingMap, position Point2D) []Point2D {
	height := hikingMap.Map[position.ToKey()]
	directions := make([]Point2D, 0)

	// North
	north := Point2D{X: position.X, Y: position.Y - 1}
	if hikingMap.Map[north.ToKey()] == height+1 {
		directions = append(directions, north)
	}

	// East
	east := Point2D{X: position.X + 1, Y: position.Y}
	if hikingMap.Map[east.ToKey()] == height+1 {
		directions = append(directions, east)
	}

	// South
	south := Point2D{X: position.X, Y: position.Y + 1}
	if hikingMap.Map[south.ToKey()] == height+1 {
		directions = append(directions, south)
	}

	// West
	west := Point2D{X: position.X - 1, Y: position.Y}
	if hikingMap.Map[west.ToKey()] == height+1 {
		directions = append(directions, west)
	}

	return directions
}

func parseHikingMap(puzzleInput string) HikingMap {
	hikingMap := HikingMap{
		Map:            make(map[string]int),
		StartingPoints: make([]Point2D, 0),
	}

	rows := input.EachLineAsString(puzzleInput)
	for y, row := range rows {
		for x, val := range strings.Split(row, "") {
			height, err := strconv.ParseInt(val, 10, 32)
			if err != nil {
				log.Fatalf("unable parse %s to  int: %v", val, err)
			}
			key := fmt.Sprintf("%d,%d", x, y)
			hikingMap.Map[key] = int(height)

			if height == 0 {
				hikingMap.StartingPoints = append(hikingMap.StartingPoints, Point2D{X: x, Y: y})
			}
		}
	}

	return hikingMap
}
