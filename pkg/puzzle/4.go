package puzzle

import (
	"fmt"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day4 struct{}

type gridLocation struct {
	col int
	row int
}

func (d Day4) Step1(puzzleInput string) string {

	grid := input.CharacterGrid(puzzleInput)
	xs := findLetter(grid, "X")

	xmasCount := 0
	for _, xLoc := range xs {
		xmasCount += findXmasCount(grid, xLoc)
	}

	return fmt.Sprintf("%d", xmasCount)
}

func (d Day4) Step2(puzzleInput string) string {
	grid := input.CharacterGrid(puzzleInput)
	as := findLetter(grid, "A")

	masXCount := 0
	for _, aLoc := range as {
		if ismasX(grid, aLoc) {
			masXCount++
		}
	}

	return fmt.Sprintf("%d", masXCount)
}

func findLetter(grid [][]string, letter string) []gridLocation {
	result := make([]gridLocation, 0)
	for rowIndex, row := range grid {
		for colIndex, col := range row {
			if col == letter {
				result = append(result, gridLocation{row: rowIndex, col: colIndex})
			}
		}
	}
	return result
}

func findXmasCount(grid [][]string, gridLoc gridLocation) int {
	xmasCount := 0
	// North
	if gridLoc.row >= 3 {
		m := grid[gridLoc.row-1][gridLoc.col]
		a := grid[gridLoc.row-2][gridLoc.col]
		s := grid[gridLoc.row-3][gridLoc.col]
		if fmt.Sprintf("%s%s%s", m, a, s) == "MAS" {
			xmasCount++
		}
	}

	// North-East
	if gridLoc.row >= 3 && gridLoc.col <= len(grid[0])-4 {
		m := grid[gridLoc.row-1][gridLoc.col+1]
		a := grid[gridLoc.row-2][gridLoc.col+2]
		s := grid[gridLoc.row-3][gridLoc.col+3]
		if fmt.Sprintf("%s%s%s", m, a, s) == "MAS" {
			xmasCount++
		}
	}

	// East
	if gridLoc.col <= len(grid[0])-4 {
		m := grid[gridLoc.row][gridLoc.col+1]
		a := grid[gridLoc.row][gridLoc.col+2]
		s := grid[gridLoc.row][gridLoc.col+3]
		if fmt.Sprintf("%s%s%s", m, a, s) == "MAS" {
			xmasCount++
		}
	}

	// South-East
	if gridLoc.row <= len(grid)-4 && gridLoc.col <= len(grid[0])-4 {
		m := grid[gridLoc.row+1][gridLoc.col+1]
		a := grid[gridLoc.row+2][gridLoc.col+2]
		s := grid[gridLoc.row+3][gridLoc.col+3]
		if fmt.Sprintf("%s%s%s", m, a, s) == "MAS" {
			xmasCount++
		}
	}

	// South
	if gridLoc.row <= len(grid)-4 {
		m := grid[gridLoc.row+1][gridLoc.col]
		a := grid[gridLoc.row+2][gridLoc.col]
		s := grid[gridLoc.row+3][gridLoc.col]
		if fmt.Sprintf("%s%s%s", m, a, s) == "MAS" {
			xmasCount++
		}
	}

	// South-West
	if gridLoc.row <= len(grid)-4 && gridLoc.col >= 3 {
		m := grid[gridLoc.row+1][gridLoc.col-1]
		a := grid[gridLoc.row+2][gridLoc.col-2]
		s := grid[gridLoc.row+3][gridLoc.col-3]
		if fmt.Sprintf("%s%s%s", m, a, s) == "MAS" {
			xmasCount++
		}
	}

	// West
	if gridLoc.col >= 3 {
		m := grid[gridLoc.row][gridLoc.col-1]
		a := grid[gridLoc.row][gridLoc.col-2]
		s := grid[gridLoc.row][gridLoc.col-3]
		if fmt.Sprintf("%s%s%s", m, a, s) == "MAS" {
			xmasCount++
		}
	}

	// North-West
	if gridLoc.row >= 3 && gridLoc.col >= 3 {
		m := grid[gridLoc.row-1][gridLoc.col-1]
		a := grid[gridLoc.row-2][gridLoc.col-2]
		s := grid[gridLoc.row-3][gridLoc.col-3]
		if fmt.Sprintf("%s%s%s", m, a, s) == "MAS" {
			xmasCount++
		}
	}

	return xmasCount
}

func ismasX(grid [][]string, gridLoc gridLocation) bool {
	// Ignore side rows because we can't have X shape words there
	if gridLoc.col == 0 || gridLoc.col == len(grid)-1 || gridLoc.row == 0 || gridLoc.row == len(grid[0])-1 {
		return false
	}

	topRight := grid[gridLoc.row-1][gridLoc.col+1]
	bottomRight := grid[gridLoc.row+1][gridLoc.col+1]
	bottomLeft := grid[gridLoc.row+1][gridLoc.col-1]
	topLeft := grid[gridLoc.row-1][gridLoc.col-1]

	diagonal1 := fmt.Sprintf("%s%s", bottomLeft, topRight)
	diagonal2 := fmt.Sprintf("%s%s", bottomRight, topLeft)

	return (diagonal1 == "MS" || diagonal1 == "SM") && (diagonal2 == "MS" || diagonal2 == "SM")
}
