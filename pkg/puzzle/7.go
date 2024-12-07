package puzzle

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day7 struct{}

type Equation struct {
	TestValue int64
	Numbers   []int64
}

type EquationPart struct {
	Number    int64
	Operation string
}

func (d Day7) Step1(puzzleInput string) string {
	validSum := validEquationsSum(puzzleInput, []string{"+", "*"})
	return fmt.Sprintf("%d", validSum)
}

func (d Day7) Step2(puzzleInput string) string {
	validSum := validEquationsSum(puzzleInput, []string{"+", "*", "||"})
	return fmt.Sprintf("%d", validSum)
}

func parseEquations(puzzleInput string) []Equation {
	equations := make([]Equation, 0)
	lines := input.EachLineAsString(puzzleInput)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		testValue, err := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		numbers := make([]int64, 0)
		rightParts := strings.Split(strings.TrimSpace(parts[1]), " ")
		for _, part := range rightParts {
			number, err := strconv.ParseInt(strings.TrimSpace(part), 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			numbers = append(numbers, int64(number))
		}
		equations = append(equations, Equation{TestValue: testValue, Numbers: numbers})
	}
	return equations
}

func validEquationsSum(puzzleInput string, operations []string) int64 {
	equations := parseEquations(puzzleInput)
	validEquations := make([]Equation, 0)
	for _, equation := range equations {
		if IsEquationValid(equation.TestValue, operations, 0, equation.Numbers) {
			validEquations = append(validEquations, equation)
		}
	}

	var validSum int64 = 0
	for _, equation := range validEquations {
		validSum += equation.TestValue
	}

	return validSum
}

func IsEquationValid(testValue int64, operations []string, left int64, right []int64) bool {
	// If right side has just one number, we can't go any deeper and
	// we can check if formula is valid.
	if len(right) == 1 {
		for _, op := range operations {
			if calculateOp(left, right[0], op) == testValue {
				return true
			}
		}
		return false
	}

	// Otherwise we need to go deeper and call this function again
	// with both operators. Move one number from right to left and
	// call the function again
	number, newRight := popLeftNumber(right)
	for _, op := range operations {
		newLeft := calculateOp(left, number, op)

		// If left side is bigger then testValue, we can stop
		// checking because the end result is guaranteed to be
		// too big
		if newLeft > testValue {
			return false
		}

		if IsEquationValid(testValue, operations, newLeft, newRight) {
			return true
		}
	}

	return false
}

func calculateOp(x, y int64, op string) int64 {
	if op == "+" {
		return x + y
	} else if op == "*" {
		return x * y
	} else if op == "||" {
		concatenated := fmt.Sprintf("%d%d", x, y)
		result, err := strconv.ParseInt(concatenated, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		return result
	}

	log.Fatal("Invalid operation")
	return 0

}

func popLeftNumber(right []int64) (int64, []int64) {
	number := right[:1][0]
	newRight := right[1:]

	return number, newRight
}
