package puzzle

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/saaste/advent-of-code-2024/pkg/input"
)

type Day11 struct{}

func (d Day11) Step1(puzzleInput string) string {
	stones := input.SpaceSeparatedInts(puzzleInput)

	stoneCount := int64(0)
	valueCache := make(map[string]int64)
	for _, stone := range stones {
		stoneCount += blink(stone, 1, 0, 25, valueCache)
	}
	return fmt.Sprintf("%d", stoneCount)
}

func (d Day11) Step2(puzzleInput string) string {
	stones := input.SpaceSeparatedInts(puzzleInput)

	stoneCount := int64(0)
	valueCache := make(map[string]int64)
	for _, stone := range stones {
		stoneCount += blink(stone, 1, 0, 75, valueCache)
	}
	return fmt.Sprintf("%d", stoneCount)
}

func blink(value int64, stoneCount int64, blinkCount int, targetBlinks int, valueCache map[string]int64) int64 {
	cacheKey := createCacheKey(value, blinkCount)

	if cachedValue, found := valueCache[cacheKey]; found {
		// If we have cached value, we can just return it
		return cachedValue
	}

	if blinkCount == targetBlinks {
		// We've reached maximum blinks!
		// Cache the stone count and return it
		valueCache[cacheKey] = stoneCount
		return stoneCount
	}

	if value == 0 {
		// Rule 1: 0 becomes 1
		// Call the function again with 1. Cache and return the result
		result := blink(1, stoneCount, blinkCount+1, targetBlinks, valueCache)
		valueCache[cacheKey] = result
		return result
	}

	// Number with event digits becomes two
	valueString := fmt.Sprintf("%d", value)
	if len(valueString)%2 == 0 {
		// Rule 2: number with event digits is split into two
		// Call the function with both numbers. Cache and return the sum of the results
		a, b := splitNumber(valueString)
		aResult := blink(a, stoneCount, blinkCount+1, targetBlinks, valueCache)
		bResult := blink(b, stoneCount, blinkCount+1, targetBlinks, valueCache)

		result := aResult + bResult
		valueCache[cacheKey] = result
		return result
	}

	// Rule 3: Multiply by 2024
	result := blink(value*2024, stoneCount, blinkCount+1, targetBlinks, valueCache)
	valueCache[cacheKey] = result

	return result
}

func splitNumber(value string) (int64, int64) {
	valueLength := len(value)
	// Split even digits stones into two stones
	left := value[:int32(valueLength)/2]
	right := value[valueLength/2:]

	leftNumber, err := strconv.ParseInt(left, 10, 64)
	if err != nil {
		log.Fatalf("unable to parse %s to int: %v", left, err)
	}

	right = strings.TrimPrefix(right, "0")
	if right == "" {
		right = "0"
	}
	rightNumber, err := strconv.ParseInt(right, 10, 64)
	if err != nil {
		log.Fatalf("unable to parse %s to int: %v", right, err)
	}

	if leftNumber < 0 || rightNumber < 0 {
		log.Fatalln("negative")
	}

	return leftNumber, rightNumber
}

func createCacheKey(value int64, blinkCount int) string {
	return fmt.Sprintf("%d,%d", value, blinkCount)
}
