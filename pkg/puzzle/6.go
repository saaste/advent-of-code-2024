package puzzle

import (
	"fmt"
	"strings"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day6 struct{}

type Point2D struct {
	X int
	Y int
}

func (p *Point2D) Equals(p2 *Point2D) bool {
	return p.X == p2.X && p.Y == p2.Y
}

func (p *Point2D) ToKey() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

type Point2DWithFacing struct {
	X      int
	Y      int
	Vector Point2D
}

type MoveActionResult struct {
	UniqueLocationCount   int
	VisitedLocations      []Point2D
	StepsUntilOutOfBounds int
}

func (d Day6) Step1(puzzleInput string) string {
	guardMap := parseMapIntoGrid(puzzleInput)
	result := moveUntilOutOfBounds(guardMap)
	return fmt.Sprintf("%d", result.UniqueLocationCount)
}

func (d Day6) Step2(puzzleInput string) string {

	guardMap := parseMapIntoGrid(puzzleInput)
	startingPosition, _ := findStartingPosition(guardMap)
	initialWalkResult := moveUntilOutOfBounds(guardMap)
	locationsCausingLoopCount := 0

	// Try to block all locations from the initial walk to see which ones
	// lead to circular path
	for _, visitedLocation := range initialWalkResult.VisitedLocations {

		// Collection for all location and facing
		visitedLocationsWithFacing := make(map[string]Point2DWithFacing)
		currentPosition := &Point2D{X: startingPosition.X, Y: startingPosition.Y}
		movingVector := &Point2D{X: 0, Y: -1}
		outOfBounds := false

		// Skip starting location. We can't put a wall where the guard is standing
		if visitedLocation.Equals(startingPosition) {
			continue
		}

		// Store original content and mark the position as a wall
		originalContent := guardMap[visitedLocation.Y][visitedLocation.X]
		guardMap[visitedLocation.Y][visitedLocation.X] = "#"

		visitedLocations := make(map[string]bool)

		// Move until we are out of bounds
		for !outOfBounds {
			key := createPositionWithFacingKey(currentPosition.X, currentPosition.Y, *movingVector)

			// Check if guard has been on the same location facing the same direction
			if _, found := visitedLocationsWithFacing[key]; found {
				locationsCausingLoopCount++
				break
			}

			// Store visited location and facing
			visitedLocationsWithFacing[key] = Point2DWithFacing{X: currentPosition.X, Y: currentPosition.Y, Vector: *movingVector}

			// Turn guard if facing a wall
			for isBlocked(guardMap, currentPosition, movingVector) {
				movingVector = turnRight(movingVector)
			}

			visitedLocations[createPositionKey(currentPosition.X, currentPosition.Y)] = true

			// Move one step
			currentPosition = move(currentPosition, movingVector)

			// Check if we are out of bounds
			outOfBounds = isOutOfBounds(guardMap, *currentPosition)
		}

		// Undo wall set
		guardMap[visitedLocation.Y][visitedLocation.X] = originalContent
	}

	return fmt.Sprintf("%d", locationsCausingLoopCount)
}

func parseMapIntoGrid(puzzleInput string) [][]string {
	guardMap := make([][]string, 0)
	lines := input.EachLineAsString(puzzleInput)
	for _, line := range lines {
		guardMap = append(guardMap, strings.Split(strings.TrimSpace(line), ""))
	}
	return guardMap
}

func findStartingPosition(guardMap [][]string) (*Point2D, error) {
	for y, row := range guardMap {
		for x, col := range row {
			if col == "^" {
				return &Point2D{X: x, Y: y}, nil
			}
		}
	}
	return nil, fmt.Errorf("unable to find starting position")
}

func isBlocked(guardMap [][]string, currentPosition *Point2D, movingVector *Point2D) bool {
	nextX := currentPosition.X + movingVector.X
	nextY := currentPosition.Y + movingVector.Y
	if nextX > len(guardMap[0])-1 || nextY > len(guardMap)-1 {
		return false
	}

	if nextX < 0 || nextY < 0 {
		return false
	}

	return guardMap[nextY][nextX] == "#"
}

func move(currentPosition *Point2D, movingVector *Point2D) *Point2D {
	return &Point2D{
		X: currentPosition.X + movingVector.X,
		Y: currentPosition.Y + movingVector.Y,
	}
}

func turnRight(movingVector *Point2D) *Point2D {
	newX := -movingVector.Y
	newY := movingVector.X
	return &Point2D{X: newX, Y: newY}
}

func isOutOfBounds(guardMap [][]string, currentLocation Point2D) bool {
	if currentLocation.X < 0 || currentLocation.Y < 0 {
		return true
	}

	if currentLocation.X > len(guardMap[0])-1 || currentLocation.Y > len(guardMap)-1 {
		return true
	}

	return false
}

func createPositionKey(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func createPositionWithFacingKey(x, y int, vector Point2D) string {
	return fmt.Sprintf("%d,%d,%d,%d", x, y, vector.X, vector.Y)
}

func moveUntilOutOfBounds(guardMap [][]string) *MoveActionResult {
	currentPosition, _ := findStartingPosition(guardMap)
	movingVector := &Point2D{X: 0, Y: -1}
	outOfBounds := false
	stepsUntilOutOfBounds := 0
	visitedLocations := make(map[string]Point2D)

	// Move until we are out of bounds
	for !outOfBounds {
		visitedLocations[createPositionKey(currentPosition.X, currentPosition.Y)] = *currentPosition

		// Turn guard if facing a wall
		if isBlocked(guardMap, currentPosition, movingVector) {
			movingVector = turnRight(movingVector)
		}

		// Move one step
		currentPosition = move(currentPosition, movingVector)

		// Check if we are out of bounds
		outOfBounds = isOutOfBounds(guardMap, *currentPosition)

		// Increase the step counter
		stepsUntilOutOfBounds++
	}

	visitedLocationPoints := make([]Point2D, 0)
	for _, point := range visitedLocations {
		visitedLocationPoints = append(visitedLocationPoints, point)
	}

	return &MoveActionResult{
		UniqueLocationCount:   len(visitedLocations),
		VisitedLocations:      visitedLocationPoints,
		StepsUntilOutOfBounds: stepsUntilOutOfBounds,
	}
}
