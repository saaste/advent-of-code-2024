package puzzle

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day13 struct{}

type Point2DInt64 struct {
	X int64
	Y int64
}

type Machine struct {
	MoveA Point2DInt64
	MoveB Point2DInt64
	Prize Point2DInt64
}

func (d Day13) Step1(puzzleInput string) string {
	machines := parseMachines(puzzleInput, false)
	tokensSpent := int64(0)
	for _, machine := range machines {
		tokensSpent += cramersRule(machine)
	}
	return fmt.Sprintf("%d", tokensSpent)
}

func (d Day13) Step2(puzzleInput string) string {
	machines := parseMachines(puzzleInput, true)

	totalPrice := int64(0)
	for _, machine := range machines {
		totalPrice += cramersRule(machine)

	}

	return fmt.Sprintf("%d", totalPrice)
}

func cramersRule(machine Machine) int64 {
	determinant := machine.MoveA.X*machine.MoveB.Y - machine.MoveA.Y*machine.MoveB.X
	n := (machine.Prize.X*machine.MoveB.Y - machine.Prize.Y*machine.MoveB.X) / determinant
	m := (machine.MoveA.X*machine.Prize.Y - machine.MoveA.Y*machine.Prize.X) / determinant

	if machine.MoveA.X*n+machine.MoveB.X*m == machine.Prize.X && machine.MoveA.Y*n+machine.MoveB.Y*m == machine.Prize.Y {
		return n*3 + m
	}

	return 0
}

func parseMachines(puzzleInput string, makeItBig bool) []Machine {
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

			x, err := strconv.ParseInt(part1, 10, 64)
			if err != nil {
				log.Fatalf("unable to parse %s to int", part1)
			}

			y, err := strconv.ParseInt(part2, 10, 64)
			if err != nil {
				log.Fatalf("unable to parse %s to int", part2)
			}

			currentMachine.MoveA = Point2DInt64{X: x, Y: y}
			continue
		}

		if strings.HasPrefix(line, "Button B") {
			parts := strings.Split(line, ", ")
			part1 := strings.ReplaceAll(parts[0], "Button B: X", "")
			part2 := strings.ReplaceAll(parts[1], "Y+", "")

			x, err := strconv.ParseInt(part1, 10, 64)
			if err != nil {
				log.Fatalf("unable to parse %s to int", part1)
			}

			y, err := strconv.ParseInt(part2, 10, 64)
			if err != nil {
				log.Fatalf("unable to parse %s to int", part2)
			}

			currentMachine.MoveB = Point2DInt64{X: x, Y: y}
			continue
		}

		if strings.HasPrefix(line, "Prize") {
			parts := strings.Split(line, ", ")
			part1 := strings.ReplaceAll(parts[0], "Prize: X=", "")
			part2 := strings.ReplaceAll(parts[1], "Y=", "")

			x, err := strconv.ParseInt(part1, 10, 64)
			if err != nil {
				log.Fatalf("unable to parse %s to int", part1)
			}

			y, err := strconv.ParseInt(part2, 10, 64)
			if err != nil {
				log.Fatalf("unable to parse %s to int", part2)
			}

			if makeItBig {
				x = x + 10000000000000
				y = y + 10000000000000
			}

			currentMachine.Prize = Point2DInt64{X: x, Y: y}
			machines = append(machines, currentMachine)
			continue
		}

		continue
	}

	return machines
}
