package puzzle

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Day3 struct{}

func (d Day3) Step1(puzzleInput string) string {
	sum := multiply(puzzleInput)
	return fmt.Sprintf("%d", sum)
}

func (d Day3) Step2(puzzleInput string) string {
	puzzleInput = removeDontParts(puzzleInput)
	sum := multiply(puzzleInput)
	return fmt.Sprintf("%d", sum)
}

// Find multiply-operations and sum the results
func multiply(puzzleInput string) int {
	r := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))`)
	matches := r.FindAllString(puzzleInput, -1)

	sum := 0
	for _, match := range matches {
		match = strings.ReplaceAll(match, "mul(", "")
		match = strings.ReplaceAll(match, ")", "")
		parts := strings.Split(match, ",")

		a, err := strconv.ParseInt(parts[0], 10, 32)
		if err != nil {
			log.Fatalf("unable to parse int a: %s - %v", parts[0], err)
		}

		b, err := strconv.ParseInt(parts[1], 10, 32)
		if err != nil {
			log.Fatalf("unable to parse int b: %s - %v", parts[1], err)
		}

		sum += int(a) * int(b)
	}
	return sum
}

// Recursive function that removes don't() parts from the input
func removeDontParts(puzzleInput string) string {
	dontIndex := strings.Index(puzzleInput, "don't()")
	if dontIndex == -1 {
		return puzzleInput
	}

	substring := puzzleInput[dontIndex:]

	doIndex := strings.Index(substring, "do()")
	if doIndex == -1 {
		return puzzleInput[:dontIndex]
	}

	substring = substring[:doIndex]
	puzzleInput = strings.ReplaceAll(puzzleInput, substring, "")

	return removeDontParts(puzzleInput)
}
