package puzzle

import (
	"fmt"
	"strings"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day8 struct{}

func (d Day8) Step1(puzzleInput string) string {
	puzzleLength := len(input.EachLineAsString(puzzleInput))
	antennas := parseAntennas(puzzleInput)
	uniqueAntiNodes := collectUniqueAntiNodes(antennas, puzzleLength)
	return fmt.Sprintf("%d", len(uniqueAntiNodes))
}

func (d Day8) Step2(puzzleInput string) string {
	puzzleLength := len(input.EachLineAsString(puzzleInput))
	antennas := parseAntennas(puzzleInput)
	uniqueAntiNodes := collectUniqueAntiNodesForever(antennas, puzzleLength)
	return fmt.Sprintf("%d", len(uniqueAntiNodes))
}

func parseAntennas(puzzleInput string) map[string][]Point2D {
	antennas := make(map[string][]Point2D)

	lines := input.EachLineAsString(puzzleInput)
	for y, line := range lines {
		for x, char := range strings.Split(line, "") {
			if char != "." {
				if _, found := antennas[char]; !found {
					antennas[char] = make([]Point2D, 0)
				}
				antennas[char] = append(antennas[char], Point2D{X: x, Y: y})
			}
		}
	}

	return antennas
}

func collectUniqueAntiNodes(antennas map[string][]Point2D, gridSize int) map[string]Point2D {
	antiNodeSet := make(map[string]Point2D)

	for _, antennaGroup := range antennas {
		for i := 0; i < len(antennaGroup)-1; i++ {
			for j := i + 1; j < len(antennaGroup); j++ {
				antennaA := antennaGroup[i]
				antennaB := antennaGroup[j]
				vectorA := Point2D{X: antennaA.X - antennaB.X, Y: antennaA.Y - antennaB.Y}
				vectorB := Point2D{X: antennaB.X - antennaA.X, Y: antennaB.Y - antennaA.Y}
				antiNodeA := Point2D{X: antennaA.X + vectorA.X, Y: antennaA.Y + vectorA.Y}
				antiNodeB := Point2D{X: antennaB.X + vectorB.X, Y: antennaB.Y + vectorB.Y}
				if isInsideGrid(gridSize, antiNodeA) {
					antiNodeSet[antiNodeA.ToKey()] = antiNodeA
				}

				if isInsideGrid(gridSize, antiNodeB) {
					antiNodeSet[antiNodeB.ToKey()] = antiNodeB
				}
			}
		}
	}

	return antiNodeSet
}

func collectUniqueAntiNodesForever(antennas map[string][]Point2D, gridSize int) map[string]Point2D {
	antiNodeSet := make(map[string]Point2D)

	for _, antennaGroup := range antennas {
		for i := 0; i < len(antennaGroup)-1; i++ {
			for j := i + 1; j < len(antennaGroup); j++ {
				antennaA := antennaGroup[i]
				antennaB := antennaGroup[j]

				antiNodeSet[antennaA.ToKey()] = antennaA
				antiNodeSet[antennaB.ToKey()] = antennaB

				vectorA := Point2D{X: antennaA.X - antennaB.X, Y: antennaA.Y - antennaB.Y}
				vectorB := Point2D{X: antennaB.X - antennaA.X, Y: antennaB.Y - antennaA.Y}

				// Generate antinodes for direction A until we are out of bounds
				antiNodeA := Point2D{X: antennaA.X + vectorA.X, Y: antennaA.Y + vectorA.Y}
				for isInsideGrid(gridSize, antiNodeA) {
					antiNodeSet[antiNodeA.ToKey()] = antiNodeA
					antiNodeA = Point2D{X: antiNodeA.X + vectorA.X, Y: antiNodeA.Y + vectorA.Y}
				}

				// Generate antinodes for direction B until we are out of bounds
				antiNodeB := Point2D{X: antennaB.X + vectorB.X, Y: antennaB.Y + vectorB.Y}
				for isInsideGrid(gridSize, antiNodeB) {
					antiNodeSet[antiNodeB.ToKey()] = antiNodeB
					antiNodeB = Point2D{X: antiNodeB.X + vectorB.X, Y: antiNodeB.Y + vectorB.Y}
				}
			}
		}
	}

	return antiNodeSet
}

func isInsideGrid(gridSize int, point Point2D) bool {
	if point.X >= 0 && point.X < gridSize && point.Y >= 0 && point.Y < gridSize {
		return true
	}
	return false
}
