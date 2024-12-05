package puzzle

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day5 struct{}

func (d Day5) Step1(puzzleInput string) string {
	rules := parseRules(puzzleInput)
	updates := parseUpdates(puzzleInput)

	validUpdates := make([][]int, 0)

	// Collect valid rules
	for _, update := range updates {
		if isValidUpdate(rules, update) {
			validUpdates = append(validUpdates, update)
		}
	}

	sum := 0
	// Sum middle element from each update
	for _, validUpdate := range validUpdates {
		sum += findMiddleElement(validUpdate)
	}

	return fmt.Sprintf("%d", sum)
}

func (d Day5) Step2(puzzleInput string) string {
	rules := parseRules(puzzleInput)
	updates := parseUpdates(puzzleInput)

	invalidUpdates := make([][]int, 0)

	// Collect invalid rules
	for _, update := range updates {
		if !isValidUpdate(rules, update) {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	// Collect fixed updates
	fixedUpdates := make([][]int, 0)
	for _, update := range invalidUpdates {
		fixedUpdates = append(fixedUpdates, fixInvalidUpdate(rules, update))
	}

	sum := 0
	// Sum middle element from each update
	for _, validUpdate := range fixedUpdates {
		sum += findMiddleElement(validUpdate)
	}

	return fmt.Sprintf("%d", sum)
}

// Parse rules into map
func parseRules(puzzleInput string) map[string][]int {
	parts := input.SplitByEmptyLine(puzzleInput)
	rules := make(map[string][]int, 0)
	ruleStrings := input.EachLineAsString(parts[0])
	for _, ruleString := range ruleStrings {
		ruleParts := strings.Split(ruleString, "|")
		x, err := strconv.ParseInt(ruleParts[0], 10, 32)
		if err != nil {
			log.Fatalf("failed to parse X from %s", ruleParts[0])
		}

		y, err := strconv.ParseInt(ruleParts[1], 10, 32)
		if err != nil {
			log.Fatalf("failed to parse Y from %s", ruleParts[1])
		}

		key := fmt.Sprintf("%d,%d", min(x, y), max(x, y))
		value := []int{int(x), int(y)}
		rules[key] = value
	}
	return rules
}

// Parse updates into collection of int slices
func parseUpdates(puzzleInput string) [][]int {
	parts := input.SplitByEmptyLine(puzzleInput)
	updates := make([][]int, 0)
	updateStrings := input.EachLineAsString(parts[1])
	for _, updateString := range updateStrings {
		update := make([]int, 0)
		pageNumbers := strings.Split(updateString, ",")
		for _, number := range pageNumbers {
			n, err := strconv.ParseInt(number, 10, 32)
			if err != nil {
				log.Fatalf("failed to parse int from %s", number)
			}
			update = append(update, int(n))
		}
		updates = append(updates, update)
	}
	return updates
}

// Validate an update against rules
func isValidUpdate(rules map[string][]int, update []int) bool {
	// Check each number against every other number after it
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			x := update[i]
			y := update[j]

			// If number pair is ok, continue to the next number pair
			if isValidPagePair(rules, x, y) {
				continue
			}

			// Order is invalid which means update is invalid
			return false
		}
	}

	// None of the rules is invalid so update must be valid
	return true
}

// Find the middle element from the update
func findMiddleElement(update []int) int {
	middleIndex := (len(update) - 1) / 2
	return update[middleIndex]
}

// Recursive function that fixes page pairs until update is valid
func fixInvalidUpdate(rules map[string][]int, update []int) []int {
	// Check each number against every other number after it
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			x := update[i]
			y := update[j]

			// If number pair is ok, continue to the next number pair
			if isValidPagePair(rules, x, y) {
				continue
			}

			// Order is invalid, switch the numbers and
			// run the fix again with the new update
			update[i] = y
			update[j] = x
			return fixInvalidUpdate(rules, update)
		}
	}

	// None of the rules is invalid so update is now valid
	return update
}

// Validate a number pair against the rules
func isValidPagePair(rules map[string][]int, x, y int) bool {
	// Find a matching rule
	key := fmt.Sprintf("%d,%d", min(x, y), max(x, y))
	rule, found := rules[key]

	// Rule not found, all ok
	if !found {
		return true
	}

	// Numbers match the order, all ok, all OK
	if x == rule[0] {
		return true
	}

	// Otherwise the numbers are wrong
	return false
}
