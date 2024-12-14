package puzzle

import (
	"fmt"
	"strings"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day14 struct{}

type Robot struct {
	OrigPos  Point2DInt64
	Pos      Point2DInt64
	Velocity Point2DInt64
}

func (d Day14) Step1(puzzleInput string) string {
	robots := parseRobots(puzzleInput)

	gridWidth := 101
	gridHeight := 103

	moveRobots(100, gridWidth, gridHeight, robots)

	tl, tr, bl, br := calculateRobotsInQuadrants(gridWidth, gridHeight, robots)

	return fmt.Sprintf("%d", tl*tr*bl*br)
}

func (d Day14) Step2(puzzleInput string) string {
	robots := parseRobots(puzzleInput)

	gridWidth := 101
	gridHeight := 103

	// Put robot positions into map for faster checking
	posMap := make(map[string]int)
	for _, robot := range robots {
		posKey := robot.Pos.GetKey()
		posMap[posKey]++
	}

	christmastTreeFoundIn := 0

	// Try to loop the first 10 000 seconds
mainLoop:
	for i := 1; i < 10000; i++ {

		// Move robot by one second
		moveRobotsAndUpdateMap(1, gridWidth, gridHeight, robots, posMap)

		looksLikeATree := false

		// Loop through each robot and check if they form a group of robots
		// next to each other
		for _, robot := range robots {

			// Surrounding areas
			neighbors := []Point2DInt64{
				{X: robot.Pos.X + 1, Y: robot.OrigPos.Y},
				{X: robot.Pos.X - 1, Y: robot.OrigPos.Y},
				{X: robot.Pos.X, Y: robot.OrigPos.Y + 1},
				{X: robot.Pos.X, Y: robot.OrigPos.Y + 1},
				{X: robot.Pos.X + 1, Y: robot.OrigPos.Y + 1},
				{X: robot.Pos.X + 1, Y: robot.OrigPos.Y - 1},
				{X: robot.Pos.X - 1, Y: robot.OrigPos.Y + 1},
				{X: robot.Pos.X - 1, Y: robot.OrigPos.Y - 1},
				{X: robot.Pos.X + 2, Y: robot.OrigPos.Y},
				{X: robot.Pos.X - 2, Y: robot.OrigPos.Y},
				{X: robot.Pos.X, Y: robot.OrigPos.Y + 2},
				{X: robot.Pos.X, Y: robot.OrigPos.Y + 2},
				{X: robot.Pos.X + 3, Y: robot.OrigPos.Y},
				{X: robot.Pos.X - 3, Y: robot.OrigPos.Y},
				{X: robot.Pos.X, Y: robot.OrigPos.Y + 3},
				{X: robot.Pos.X, Y: robot.OrigPos.Y + 3},
			}

			if !looksLikeATree {
				matches := 0
				for _, neighbor := range neighbors {
					if roboCount, found := posMap[neighbor.GetKey()]; found && roboCount > 0 {
						matches++
					} else {
						// If even one of the surrounding squares are empty, this
						// is not the potential tree position we are looking for
						break
					}
				}

				// All surrounding squares have drones! I think we have a christmas tree here! Exciting!
				if matches == len(neighbors) {
					christmastTreeFoundIn = i
					break mainLoop
				}
			}
		}
	}

	return fmt.Sprintf("%d", christmastTreeFoundIn)
}

func calculateRobotsInQuadrants(gridWidth, gridHeight int, robots []*Robot) (tl, tr, bl, br int) {
	// Find top left quadrant
	for y := 0; y < gridHeight/2; y++ {
		for x := 0; x < gridWidth/2; x++ {
			for _, robot := range robots {
				if robot.Pos.X == int64(x) && robot.Pos.Y == int64(y) {
					tl++
				}
			}
		}
	}

	// Find top right quadrant
	for y := 0; y < gridHeight/2; y++ {
		for x := gridWidth/2 + 1; x < gridWidth; x++ {
			for _, robot := range robots {
				if robot.Pos.X == int64(x) && robot.Pos.Y == int64(y) {
					tr++
				}
			}
		}
	}

	// Find bottom right quadrant
	for y := gridHeight/2 + 1; y < gridHeight; y++ {
		for x := gridWidth/2 + 1; x < gridWidth; x++ {
			for _, robot := range robots {
				if robot.Pos.X == int64(x) && robot.Pos.Y == int64(y) {
					br++
				}
			}
		}
	}

	// Find bottom left quadrant
	for y := gridHeight/2 + 1; y < gridHeight; y++ {
		for x := 0; x < gridWidth/2; x++ {
			for _, robot := range robots {
				if robot.Pos.X == int64(x) && robot.Pos.Y == int64(y) {
					bl++
				}
			}
		}
	}

	return
}

func moveRobots(seconds, gridWidth, gridHeight int, robots []*Robot) {
	for _, robot := range robots {
		newX := robot.Pos.X + robot.Velocity.X*int64(seconds)
		if newX >= int64(gridWidth) {
			newX = newX % int64(gridWidth)
		} else if newX < 0 {
			newMod := newX % int64(gridWidth)
			if newMod < 0 { // TODO: SELVITÄ MIKSI TÄMÄ
				newX = int64(gridWidth) + newX%int64(gridWidth)
			} else {
				newX = 0
			}

		}
		robot.Pos.X = newX

		newY := robot.Pos.Y + robot.Velocity.Y*int64(seconds)
		if newY >= int64(gridHeight) {
			newY = newY % int64(gridHeight)
		} else if newY < 0 {
			newMod := newY % int64(gridHeight)
			if newMod < 0 { // TODO: SELVITÄ MIKSI TÄMÄ
				newY = int64(gridHeight) + newY%int64(gridHeight)
			} else {
				newY = 0
			}
		}
		robot.Pos.Y = newY
	}
}

func moveRobotsAndUpdateMap(seconds, gridWidth, gridHeight int, robots []*Robot, posMap map[string]int) {
	for _, robot := range robots {
		posKey := robot.Pos.GetKey()
		posMap[posKey]--

		newX := robot.Pos.X + robot.Velocity.X*int64(seconds)
		if newX >= int64(gridWidth) {
			newX = newX % int64(gridWidth)
		} else if newX < 0 {
			newMod := newX % int64(gridWidth)
			if newMod < 0 {
				newX = int64(gridWidth) + newX%int64(gridWidth)
			} else {
				newX = 0
			}

		}
		robot.Pos.X = newX

		newY := robot.Pos.Y + robot.Velocity.Y*int64(seconds)
		if newY >= int64(gridHeight) {
			newY = newY % int64(gridHeight)
		} else if newY < 0 {
			newMod := newY % int64(gridHeight)
			if newMod < 0 {
				newY = int64(gridHeight) + newY%int64(gridHeight)
			} else {
				newY = 0
			}
		}
		robot.Pos.Y = newY

		posKey = robot.Pos.GetKey()
		posMap[posKey]++
		if posMap[posKey] == 0 {
			delete(posMap, posKey)
		}
	}
}

func parseRobots(puzzleInput string) []*Robot {
	robots := make([]*Robot, 0)

	lines := input.EachLineAsString(puzzleInput)
	for _, line := range lines {
		parts := strings.Split(line, " ")

		startParts := strings.Split(strings.ReplaceAll(parts[0], "p=", ""), ",")
		startX := input.StringAsInt64(startParts[0])
		startY := input.StringAsInt64(startParts[1])

		velocityParts := strings.Split(strings.ReplaceAll(parts[1], "v=", ""), ",")
		velocityX := input.StringAsInt64(velocityParts[0])
		velocityY := input.StringAsInt64(velocityParts[1])

		robot := &Robot{
			OrigPos:  Point2DInt64{X: startX, Y: startY},
			Pos:      Point2DInt64{X: startX, Y: startY},
			Velocity: Point2DInt64{X: velocityX, Y: velocityY}}

		robots = append(robots, robot)
	}

	return robots
}
