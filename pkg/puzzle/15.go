package puzzle

import (
	"fmt"
	"log"
	"slices"
	"sort"
	"strings"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day15 struct{}

func (d Day15) Step1(puzzleInput string) string {
	width := 50
	height := 50

	warehouse, movements, robotPos := parseWarehouseAndRobotMoves(puzzleInput)

	for _, movement := range movements {
		moved := moveWarehouseRobot(warehouse, robotPos, movement)
		if moved {
			robotPos = robotPos.Add(movement)
		}
	}

	return fmt.Sprintf("%d", calculateGPSSum(warehouse, width, height, false))
}

func (d Day15) Step2(puzzleInput string) string {
	width := 100
	height := 50

	warehouse, movements, robotPos := parseBiggerWarehouseAndRobotMoves(puzzleInput)
	for _, movement := range movements {
		moved := moveBiggerWarehouseRobot(warehouse, robotPos, movement)
		if moved {
			robotPos = robotPos.Add(movement)
		}
	}
	return fmt.Sprintf("%d", calculateGPSSum(warehouse, width, height, true))
}

func calculateGPSSum(warehouse map[string]string, width, height int, isBigger bool) int64 {
	var boxChar = "O"
	if isBigger {
		boxChar = "["
	}

	sum := int64(0)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			posKey := fmt.Sprintf("%d,%d", x, y)
			char := warehouse[posKey]
			if char == boxChar {
				sum += int64(100)*int64(y) + int64(x)
			}
		}
	}
	return sum
}

func moveWarehouseRobot(warehouse map[string]string, robotPos, movement Point2D) bool {
	canMove, boxesToMove := checkFreeSpaceAndCollectBoxes(warehouse, robotPos, movement)

	if !canMove {
		// Robot can't move, do nothing
		return false
	}

	// Move the robot
	newPosition := robotPos.Add(movement)
	warehouse[newPosition.ToKey()] = "@"
	warehouse[robotPos.ToKey()] = "."

	// Move all the boxes
	for _, boxLocation := range boxesToMove {
		newPosition := boxLocation.Add(movement)
		warehouse[newPosition.ToKey()] = "O"
	}

	return true
}

func moveBiggerWarehouseRobot(warehouse map[string]string, robotPos, movement Point2D) bool {
	canMove, boxesToMove := checkBiggerFreeSpaceAndCollectBoxes(warehouse, robotPos, movement)

	if !canMove {
		// Robot can't move, do nothing
		return false
	}

	// Sort movable boxes so that we always move the
	// outer ones first
	if movement.X != 0 {
		sort.Slice(boxesToMove, func(i, j int) bool {
			return boxesToMove[i].X < boxesToMove[j].X
		})
		if movement.X > 0 {
			slices.Reverse(boxesToMove)
		}
	} else {
		sort.Slice(boxesToMove, func(i, j int) bool {
			return boxesToMove[i].Y < boxesToMove[j].Y
		})
		if movement.Y > 0 {
			slices.Reverse(boxesToMove)
		}
	}

	// Move all the boxes
	for _, boxLocation := range boxesToMove {
		boxChar := warehouse[boxLocation.ToKey()]
		newPosition := boxLocation.Add(movement)
		warehouse[newPosition.ToKey()] = boxChar
		warehouse[boxLocation.ToKey()] = "."
	}

	// Move the robot
	newPosition := robotPos.Add(movement)
	warehouse[newPosition.ToKey()] = "@"
	warehouse[robotPos.ToKey()] = "."

	return true
}

// Check if there is space to move. If there is, return the location of
// each box that is between the robot and the empty space.
func checkFreeSpaceAndCollectBoxes(warehouse map[string]string, robotPos, movement Point2D) (bool, []Point2D) {
	boxes := make([]Point2D, 0)

	currentPosition := robotPos
	for {
		nextCell := currentPosition.Add(movement)
		nextChar, found := warehouse[nextCell.ToKey()]
		if !found {
			log.Fatalf("No cell information in the map! WTF?!")
		}

		if nextChar == "#" {
			// We hit a wall. No space to move
			return false, []Point2D{}
		}

		if nextChar == "O" {
			// We hit a box. Add to the collection
			boxes = append(boxes, nextCell)

			// Continue checking from the box location
			currentPosition = nextCell
			continue
		}

		// Otherwise it must be an empty cell so we can move
		return true, boxes
	}
}

func checkBiggerFreeSpaceAndCollectBoxes(warehouse map[string]string, robotPos, movement Point2D) (bool, []Point2D) {
	boxes := make([]Point2D, 0)

	currentPosition := robotPos
	for {
		nextCell := currentPosition.Add(movement)
		nextChar, found := warehouse[nextCell.ToKey()]
		if !found {
			log.Fatalf("No cell information in the map! WTF?!")
		}

		if nextChar == "#" {
			// We hit a wall. No space to move
			return false, []Point2D{}
		}

		if nextChar == "[" {
			// We hit the left side of the box
			// Continue checking from the box location
			currentPosition = nextCell

			// Moving horizontally, we can just continue to the next cell
			// and add this one to the list of movable boxes
			if movement.Y == 0 {
				boxes = append(boxes, nextCell)
				continue
			}

		}

		if nextChar == "]" {
			// We hit the right side of the box.
			// Continue checking from the box location
			currentPosition = nextCell

			// Moving horizontally, we can just continue to the next cell
			// and add this one to the list of movable boxes
			if movement.Y == 0 {
				boxes = append(boxes, nextCell)
				continue
			}
		}

		// Special case when we move up and down and bump into a box.
		// We need to check diagonally aligned boxes if they can be
		// moved. Recursion time!
		if movement.Y != 0 && (nextChar == "[" || nextChar == "]") {
			return checkSpaceAfterBoxRecursive(warehouse, nextCell, movement, boxes)
		}

		// Otherwise it must be an empty cell so we can move
		return true, boxes
	}
}

// Recursive function that checks the space after a box. If box is next to another
// box, check that as well. Continue until we find an empty space or a wall
func checkSpaceAfterBoxRecursive(warehouse map[string]string, boxPos, movement Point2D, boxesToMove []Point2D) (bool, []Point2D) {
	var boxLeftPos, boxRightPos Point2D

	// Figure out both sides of the box
	boxChar := warehouse[boxPos.ToKey()]
	if boxChar == "[" {
		boxLeftPos = boxPos
		boxRightPos = boxLeftPos.Add(Point2D{X: 1, Y: 0})
	} else {
		boxRightPos = boxPos
		boxLeftPos = boxRightPos.Add(Point2D{X: -1, Y: 0})
	}

	// Check spaces after both sides
	spaceLeft := boxLeftPos.Add(movement)
	spaceRight := boxRightPos.Add(movement)

	leftChar := warehouse[spaceLeft.ToKey()]
	rightChar := warehouse[spaceRight.ToKey()]

	if leftChar == "#" || rightChar == "#" {
		// We hit the wall. No moving!
		return false, []Point2D{}
	}

	boxesToMove = addBoxesToMoveIfNotExist(boxesToMove, boxLeftPos, boxRightPos)

	if leftChar == "." && rightChar == "." {
		// Both cells are empty so we can move
		return true, boxesToMove
	}

	if leftChar == "[" && rightChar == "]" {
		// We have a box directly above us so
		// check that and return whatever it returns
		return checkSpaceAfterBoxRecursive(warehouse, spaceLeft, movement, boxesToMove)
	}

	// Otherwise we need to check the next box(es)

	// Start from the right side if it is not empty
	if rightChar != "." {
		canMove, newBoxes := checkSpaceAfterBoxRecursive(warehouse, spaceRight, movement, boxesToMove)
		if !canMove {
			// We hit the wall. No moving!
			return false, []Point2D{}
		}
		boxesToMove = newBoxes
	}

	// Then check the right side if it is not empty
	if leftChar != "." {
		return checkSpaceAfterBoxRecursive(warehouse, spaceLeft, movement, boxesToMove)
	}

	// If we reach here, the box can move so... MOVE!
	return true, boxesToMove
}

// Check if box is in the list
func contains(boxesToMove []Point2D, box Point2D) bool {
	for _, oldBox := range boxesToMove {
		if oldBox.Equals(&box) {
			return true
		}
	}
	return false
}

// Add box to list if it is not there already
func addBoxesToMoveIfNotExist(boxesToMove []Point2D, newBoxes ...Point2D) []Point2D {
	for _, newBox := range newBoxes {
		if !contains(boxesToMove, newBox) {
			boxesToMove = append(boxesToMove, newBox)
		}

	}
	return boxesToMove
}

// Parse warehouse map, robot movements and starting position from the input
func parseWarehouseAndRobotMoves(puzzleInput string) (map[string]string, []Point2D, Point2D) {
	warehouse := make(map[string]string)
	movements := make([]Point2D, 0)
	startingPosition := Point2D{}

	// Build the warehouse map
	lines := input.EachLineAsString(puzzleInput)
	for y, line := range lines {
		// skip the empty line
		if len(line) == 0 {
			continue
		}

		// If line starts with # it is part of the map
		if strings.HasPrefix(line, "#") {
			for x, char := range strings.Split(line, "") {
				posKey := fmt.Sprintf("%d,%d", x, y)
				if char == "@" {
					startingPosition = Point2D{X: x, Y: y}
				}
				warehouse[posKey] = char
			}
			continue
		}

		// Otherwise it is part of the robot movements
		for _, char := range strings.Split(line, "") {
			switch char {
			case "^":
				movements = append(movements, Point2D{X: 0, Y: -1})
			case ">":
				movements = append(movements, Point2D{X: 1, Y: 0})
			case "v":
				movements = append(movements, Point2D{X: 0, Y: 1})
			case "<":
				movements = append(movements, Point2D{X: -1, Y: 0})
			default:
				log.Fatalf("Invalid direction %s", char)
			}
		}
	}

	return warehouse, movements, startingPosition
}

func parseBiggerWarehouseAndRobotMoves(puzzleInput string) (map[string]string, []Point2D, Point2D) {
	warehouse := make(map[string]string)
	movements := make([]Point2D, 0)
	startingPosition := Point2D{}

	// Make it bigger
	puzzleInput = strings.ReplaceAll(puzzleInput, "#", "##")
	puzzleInput = strings.ReplaceAll(puzzleInput, "O", "[]")
	puzzleInput = strings.ReplaceAll(puzzleInput, ".", "..")
	puzzleInput = strings.ReplaceAll(puzzleInput, "@", "@.")

	// Build the warehouse map
	lines := input.EachLineAsString(puzzleInput)
	for y, line := range lines {
		// skip the empty line
		if len(line) == 0 {
			continue
		}

		// If line starts with # it is part of the map
		if strings.HasPrefix(line, "#") {
			for x, char := range strings.Split(line, "") {
				posKey := fmt.Sprintf("%d,%d", x, y)
				if char == "@" {
					startingPosition = Point2D{X: x, Y: y}
				}
				warehouse[posKey] = char
			}

			continue
		}

		// Otherwise it is part of the robot movements
		for _, char := range strings.Split(line, "") {
			switch char {
			case "^":
				movements = append(movements, Point2D{X: 0, Y: -1})
			case ">":
				movements = append(movements, Point2D{X: 1, Y: 0})
			case "v":
				movements = append(movements, Point2D{X: 0, Y: 1})
			case "<":
				movements = append(movements, Point2D{X: -1, Y: 0})
			default:
				log.Fatalf("Invalid direction %s", char)
			}
		}
	}

	return warehouse, movements, startingPosition
}
