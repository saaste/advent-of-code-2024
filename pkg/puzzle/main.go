package puzzle

import (
	"fmt"
	"log"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type DayPuzzle interface {
	Step1(input string) string
	Step2(input string) string
}

func RunPuzzle(day, step int) {
	if step < 1 || step > 2 {
		log.Fatalf("invalid step %d", step)
	}

	puzzle, ok := puzzles[day]
	if !ok {
		log.Fatalf("failed to load the puzzle struct for day %d", day)
	}

	input := input.ReadFile(day)
	var result string
	if step == 1 {
		result = puzzle.Step1(input)
	} else {
		result = puzzle.Step2(input)
	}
	fmt.Printf("Day %d / Step %d: %s\n", day, step, result)
}

func ValidatePuzzle(day, step int) {
	if step < 1 || step > 2 {
		log.Fatalf("invalid step %d", step)
	}

	puzzle, ok := puzzles[day]
	if !ok {
		log.Fatalf("failed to load the puzzle struct for day %d", day)
	}

	input := input.ReadFile(day)

	expected := correctAnswers[day].Step1

	var actual string
	if step == 1 {
		actual = puzzle.Step1(input)
	} else {
		actual = puzzle.Step2(input)
	}

	if expected != actual {
		fmt.Printf("Day %d / Step %d is INVALID\n", day, step)
		fmt.Printf("Expected %s\n", expected)
		fmt.Printf("Actual %s\n", expected)
	} else {
		fmt.Printf("Day %d / Step %d is VALID\n", day, step)
	}
}

var puzzles = map[int]DayPuzzle{
	1:  Day1{},
	2:  Day2{},
	3:  Day3{},
	4:  Day4{},
	5:  Day5{},
	6:  Day6{},
	7:  Day7{},
	8:  Day8{},
	9:  Day9{},
	10: Day10{},
	11: Day11{},
	12: Day12{},
	13: Day13{},
	14: Day14{},
	15: Day15{},
	16: Day16{},
	17: Day17{},
	18: Day18{},
	19: Day19{},
	20: Day20{},
	21: Day21{},
	22: Day22{},
	23: Day23{},
	24: Day24{},
	25: Day25{},
}
