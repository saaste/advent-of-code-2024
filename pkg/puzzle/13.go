package puzzle

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day13 struct{}

type Machine struct {
	MoveA Point2D
	MoveB Point2D
	Prize Point2D
}

func (d Day13) Step1(puzzleInput string) string {
	machines := parseMachines(puzzleInput)
	tokensSpent := 0
	for _, machine := range machines {
		solutions := make([]Point2D, 0)
		for aMoves := 1; aMoves <= 100; aMoves++ {
			for bMoves := 1; bMoves <= 100; bMoves++ {
				xMatches := aMoves*machine.MoveA.X+bMoves*machine.MoveB.X == machine.Prize.X
				yMatches := aMoves*machine.MoveA.Y+bMoves*machine.MoveB.Y == machine.Prize.Y
				if xMatches && yMatches {
					solutions = append(solutions, Point2D{aMoves, bMoves})
				}
			}
		}
		if len(solutions) == 0 {
			continue
		}
		cheapestSolutionPrice := math.MaxInt32
		for _, solution := range solutions {
			price := solution.X*3 + solution.Y*1
			if price < cheapestSolutionPrice {
				cheapestSolutionPrice = price
			}
		}
		tokensSpent += cheapestSolutionPrice
	}
	return fmt.Sprintf("%d", tokensSpent)
}

func (d Day13) Step2(puzzleInput string) string {
	return ""
}

func parseMachines(puzzleInput string) []Machine {
	machines := make([]Machine, 0)
	lines := input.EachLineAsString(puzzleInput)

	var currentMachine Machine
	for _, line := range lines {

		if strings.HasPrefix(line, "Button A") {
			// Create a new machine
			currentMachine = Machine{}
			parts := strings.Split(line, ", ")
			part1 := strings.ReplaceAll(parts[0], "Button A: X", "")
			part2 := strings.ReplaceAll(parts[1], "Y+", "")

			x, err := strconv.ParseInt(part1, 10, 32)
			if err != nil {
				log.Fatalf("unable to parse %s to int", part1)
			}

			y, err := strconv.ParseInt(part2, 10, 32)
			if err != nil {
				log.Fatalf("unable to parse %s to int", part2)
			}

			currentMachine.MoveA = Point2D{X: int(x), Y: int(y)}
			continue
		}

		if strings.HasPrefix(line, "Button B") {
			parts := strings.Split(line, ", ")
			part1 := strings.ReplaceAll(parts[0], "Button B: X", "")
			part2 := strings.ReplaceAll(parts[1], "Y+", "")

			x, err := strconv.ParseInt(part1, 10, 32)
			if err != nil {
				log.Fatalf("unable to parse %s to int", part1)
			}

			y, err := strconv.ParseInt(part2, 10, 32)
			if err != nil {
				log.Fatalf("unable to parse %s to int", part2)
			}

			currentMachine.MoveB = Point2D{X: int(x), Y: int(y)}
			continue
		}

		if strings.HasPrefix(line, "Prize") {
			parts := strings.Split(line, ", ")
			part1 := strings.ReplaceAll(parts[0], "Prize: X=", "")
			part2 := strings.ReplaceAll(parts[1], "Y=", "")

			x, err := strconv.ParseInt(part1, 10, 32)
			if err != nil {
				log.Fatalf("unable to parse %s to int", part1)
			}

			y, err := strconv.ParseInt(part2, 10, 32)
			if err != nil {
				log.Fatalf("unable to parse %s to int", part2)
			}

			currentMachine.Prize = Point2D{X: int(x), Y: int(y)}
			machines = append(machines, currentMachine)
			continue
		}

		continue
	}

	return machines
}
